package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Write(w http.ResponseWriter, payload any) {
	w.Header().Set("Content-Type", "application/json")
	raw, _ := json.Marshal(payload)
	w.Write(raw)
}

type ApiError struct {
	Code string `json:"code"`
	Err  string `json:"err"`
}

func WriteError(w http.ResponseWriter, code string, err error) {
	Write(w, ApiError{
		Code: code,
		Err:  err.Error(),
	})
}

func WriteAck(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"ack":true}`))
}

// We separate errors who need to be send because
// client can act to change the problem and other errors
// that client don't need to see or api don't show
// for security as the database informations.
// Handle error need to be send to the client.
func WriteInputError(w http.ResponseWriter, err error) {
	WriteError(w, "input_error:", err)
}

func WriteUnauthorizedError(w http.ResponseWriter, err error) {
	WriteError(w, "unauthorized_error:", err)
	fmt.Printf("unauthorized_error: %+v\n", err)
}

func WritePaginationError(w http.ResponseWriter, err error) {
	WriteError(w, "pagination_error:", err)
}

// Handle error need to be log in the api.
// The WriteError will be return to the client
// and the printf will be log for historic information api
func WriteInternalErrorMail(w http.ResponseWriter, err error) {
	WriteError(w, "mail_error:", errors.New("internal probleme to send mail"))
	fmt.Printf("mail_error: %+v\n", err)
}

func WriteDatabaseError(w http.ResponseWriter, err error) {
	WriteError(w, "database_error:", errors.New("internal problem with database"))
	fmt.Printf("database error: %+v\n", err)
}

func WritePasswordError(w http.ResponseWriter, err error) {
	WriteError(w, "password_error:", errors.New("internal probleme with password"))
	fmt.Printf("hashing_password_error: %+v\n", err)
}

// Other helpers for a readable errors handle in api
func Read(r *http.Request, payload any) (err error) {
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(raw, payload)
	if err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ValidatePassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func ParseToken(r *http.Request, secretKeyS []byte) (int64, error) {
	var zero int64

	if r.Header.Get("Authorization") == "" {
		return zero, errors.New("missing or empty authorization header")
	}

	token, err := jwt.Parse(r.Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
		return secretKeyS, nil
	})
	if err != nil {
		return zero, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return zero, errors.New("unexpected claims format")
	}

	sub, ok := claims["sub"]
	if !ok {
		return zero, errors.New("missing claim sub")
	}

	userID, ok := sub.(float64)
	if !ok {
		return zero, errors.New("expected sub to be a number")
	}

	expired_at, ok := claims["expired_at"]
	if !ok {
		return zero, errors.New("missing claim expired_at")
	}

	expirationTimeString, ok := expired_at.(string)
	if !ok {
		return zero, errors.New("expected expired_at to be a string")
	}

	expirationTime, err := time.Parse(time.RFC3339, expirationTimeString)
	if err != nil {
		return zero, errors.New("malformed expired_at")
	}

	if expirationTime.Before(time.Now()) {
		return zero, errors.New("token expired")
	}

	return int64(userID), err
}

func (s Service) checkDataToDelete(userId int64) {
	var w http.ResponseWriter
	var recipes []Recipe
	var shoppingList []int64

	// recipes
	_, err := s.DB.Exec(`
		UPDATE recipe
		SET
		 owner_id = 1
		WHERE is_public = 1
		AND owner_id = ?
	`, userId)
	if err != nil {
		WriteDatabaseError(w, err)
	}
	err = s.DB.Select(&recipes, `
		SELECT id
		FROM recipe 
		WHERE is_public=0
		AND owner_id = ?
	`, userId)
	if err != nil {
		WriteDatabaseError(w, err)
	}
	for _, recipe := range recipes {
		_, _ = s.DB.Exec(`
			DELETE FROM ustensil_recipe
			WHERE recipe.id = ?
		`, recipe.ID)
		_, _ = s.DB.Exec(`
			DELETE FROM meal_recipe
			WHERE recipe_id = ?
		`, recipe.ID)
		_, _ = s.DB.Exec(`
			DELETE FROM recipe_food
			WHERE recipe_id = ?
		`, recipe.ID)
		_, _ = s.DB.Exec(`
			DELETE FROM recipe
			WHERE id = ?
		`, recipe.ID)
	}
	//meals On CASCADE
	// handle shooping list, home, ustensil, preference.
	_, _ = s.DB.Exec(`
		DELETE FROM user_home
		JOIN home ON home.id = user_home.home_id
		WHERE owner_id = ?
	`, userId)
	_, _ = s.DB.Exec(`
		DELETE FROM home
		WHERE owner_id = ?
	`, userId)
	_, _ = s.DB.Exec(`
		DELETE FROM user_preference
		JOIN preference ON preference.id = user_preference.preference_id
		WHERE user_id = ?
	`, userId)
	_, _ = s.DB.Exec(`
		DELETE FROM user_ustensil
		JOIN ustensil ON ustensil.id = user_ustensil.ustensil_id
		WHERE user_id = ?
	`, userId)

	_ = s.DB.Select(&shoppingList, `
		SELECT id
		FROM shopping_list 
		JOIN user_shopping_list us ON us.shopping_list_id = shopping_list.id
		AND us.user_id = ?
	`, userId)
	for _, list := range shoppingList {
		_, err = s.DB.Exec(`
			DELETE FROM user_shopping_list
			WHERE user_id = ?
		`, userId)
		if err != nil {
			WriteDatabaseError(w, err)
		}
		_, err = s.DB.Exec(`
			DELETE FROM shopping_list
			WHERE id = ?
		`, list)
		if err != nil {
			WriteDatabaseError(w, err)
		}
	}
}

func (s Service) CheckHomePermission(homeID any, userID int64) error {
	// This func is only to check the junction between a home and
	// the members associated to it. It will be usefull for update a shopping list,
	// delete a shopping list and see shopping lists in a home.
	// Maybe i'll add it to every function that care about shopping list.
	if homeID == nil {
		return errors.New("you're not authorized")
	} else {
		var members []int64
		err := s.DB.Select(&members, `
			SELECT user_id
			FROM user_home
			WHERE home_id = ?
		`, homeID)
		if err != nil {
			return err
		}

		var i int = 0
		for _, member := range members {
			if member == userID {
				i += 1
			}
		}

		if i == 0 {
			return errors.New("your not a member of this home")
		}
	}

	return nil
}
