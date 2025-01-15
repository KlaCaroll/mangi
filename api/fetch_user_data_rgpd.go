package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type UserDatasOutput struct {
	ID           int64          `db:"id" json:"id"`
	Name         string         `db:"name" json:"name"`
	Password     string         `db:"password" json:"password"`
	Email        string         `db:"email" json:"email"`
	IsAdmin      bool           `db:"is_admin" json:"is_admin"`
	Homes        []Home         `json:"homes"`
	Preferences  []Preference   `json:"preferences"`
	Ustensils    []Ustensil     `json:"ustensils"`
	Recipes      []Recipe       `json:"recipes"`
	Meals        []Meal         `json:"meals"`
	ShoppingList []ShoppingList `json:"shopping_lists"`
	Comments     []Comment      `json:"comments"`
}

func (s Service) FetchUserDataRGDP(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var user UserDatasOutput
	err = s.DB.Get(&user, `
		SELECT id, name, email, password, is_admin
		FROM user
		WHERE id = ?
	`, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	err = s.DB.Select(&user.Homes, `
		SELECT id, owner_id, name
		FROM home
		JOIN user_home on user_home.home_id = home.id
		WHERE user_home.user_id = ?
	`, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	err = s.DB.Select(&user.Preferences, `
		SELECT p.name, up.exist
		FROM user_preference up
		JOIN preference p ON up.preference_id = p.id
		WHERE up.user_id = ?
		AND up.exist = 1
	`, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	err = s.DB.Select(&user.Ustensils, `
		Select u.name, uu.exist
		FROM user_ustensil uu
		JOIN ustensil u ON uu.ustensil_id = u.id
		WHERE uu.user_id = ?
	`, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	err = s.DB.Select(&user.Recipes, `
		SELECT id, name, preparation_time, total_time, description, is_public, owner_id
		FROM recipe
		WHERE owner_id = ?
	`, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}
	for i, recipe := range user.Recipes {
		// handle ingredients
		err = s.DB.Select(&user.Recipes[i].Ingredients, `
			SELECT f.id, f.name, rf.quantity, rf.unit
			FROM recipe r
			JOIN recipe_food rf ON rf.recipe_id = r.id
			JOIN food f ON f.id = rf.food_id
			WHERE r.id = ?
		`, recipe.ID)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
		// handle ustensils
		err = s.DB.Select(&user.Recipes[i].Ustensils, `
				SELECT us.name
				FROM ustensil us
				JOIN ustensil_recipe ur ON ur.ustensil_id = us.id
				WHERE ur.recipe_id = ?
			`, recipe.ID)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
		// handle categories
		err = s.DB.Select(&user.Recipes[i].Categories, `
				SELECT c.name
				FROM category c
				JOIN recipe_category rc ON rc.category_id = c.id
				WHERE rc.recipe_id = ?
			`, recipe.ID)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
	}

	err = s.DB.Select(&user.Meals, `
		SELECT m.planned_at, m.guests
		FROM meal m
		JOIN meal_recipe mr ON mr.meal_id = m.id
		WHERE m.owner_id = ?
	`, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}
	for i, meal := range user.Meals {
		err = s.DB.Select(&user.Meals[i].Recipe, `
			SELECT mr.recipe_id as id, r.name 
			FROM meal m
			JOIN meal_recipe mr ON mr.meal_id = m.id
			JOIN recipe r ON r.id = mr.recipe_id
			WHERE m.owner_id = ?
			AND m.planned_at = ?
			AND m.guests = ?
		`, userID, meal.PlannedAt, meal.Guests)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
	}

	err = s.DB.Select(&user.Comments, `
		SELECT owner_id, description, parent_id
		FROM comment
		WHERE owner_id = ?
	`, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	err = s.DB.Select(&user.ShoppingList, `
		SELECT fromTime, toTime, name, user_id
		FROM shopping_list
		WHERE user_id = ?
	`, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}
	for i, list := range user.ShoppingList {
		// handle items in each list
		err = s.DB.Select(&user.ShoppingList[i].Items, `
			SELECT food_name, food_quantity, food_unit
			FROM shopping_list
			WHERE user_id = ?
			AND fromTime = ?
			AND toTime = ?
		`, userID, list.From, list.To)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
	}

	err = s.DB.Select(&user.Comments, `
		SELECT id, owner_id, description, parent_id
		FROM comment
		WHERE owner_id = ?
	`, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	filename := fmt.Sprintf("bin/rgpd-output/output-data-%v", userID)
	file, _ := json.MarshalIndent(user, "", "")
	err = os.WriteFile(filename, file, 0644)
	if err != nil {
		WriteError(w, "internal_error", errors.New("internal_problem_with_user_data"))
		return
	}

	err = SendMailRgpd(user.Email, filename)
	if err != nil {
		WriteInternalErrorMail(w, err)
		return
	}

	WriteAck(w)
}
