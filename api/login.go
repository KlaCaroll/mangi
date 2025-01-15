package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type LoginInput struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type LoginOutput struct {
	Token string `json:"token"`
}

func (s Service) Login(w http.ResponseWriter, r *http.Request) {
	var input LoginInput
	err := Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	if (input.Password == "") || (input.Email == "") {
		WriteInputError(w, errors.New("need password and email to login"))
		return
	}

	var user struct {
		ID       int64  `db:"id" json:"id"`
		Password string `db:"password" json:"password"`
	}

	err = s.DB.Get(&user, `
		SELECT id, password
		FROM user
		WHERE email = ?
	`, input.Email)
	if err != nil {
		if user.ID == 0 {
			WriteError(w, "database_error:", errors.New("no user found"))
		} else {
			WriteDatabaseError(w, err)
		}
		return
	}

	if !ValidatePassword(input.Password, user.Password) {
		errN := errors.New("this password doesn't match")
		WriteError(w, "checking_password:", errN)
		fmt.Printf("checking password: %+v\n", err)
		return
	}

	// TODO (caroll) Rendre la durée d'expiration configurable.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":        user.ID,
		"expired_at": time.Now().Add(time.Hour * time.Duration(5)).Format(time.RFC3339),
	})

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		errN := errors.New("this password doesn't match")
		WriteError(w, "checking_password:", errN)
		fmt.Printf("signing token error: %+v\n", err)
		return
	}

	Write(w, LoginOutput{
		Token: tokenString,
	})
}
