package main

import (
	"errors"
	"net/http"
	"time"
)

type ComputeShoppingListInput struct {
	UserID int64     `json:"user_id" qs:"user_id"`
	From   time.Time `json:"from" qs:"from"`
	To     time.Time `json:"to" qs:"to"`
	Name   string    `db:"name" json:"name"`
	HomeID int64     `db:"home_id" json:"home_id,omitempty"`
}

func (s Service) ComputeShoppingList(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input ComputeShoppingListInput
	input.HomeID = 0
	err = Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	// gerer le cas admin
	if userID != input.UserID {
		WriteUnauthorizedError(w, errors.New("can't ask for the shopping list from another user"))
		return
	}

	var shoppingList ShoppingList
	err = s.DB.Select(&shoppingList.Items, `
		SELECT f.name AS food_name, SUM(rf.quantity * m.guests) AS food_quantity, MIN(rf.unit) AS food_unit
		FROM meal m
		JOIN meal_recipe mr ON m.id = mr.meal_id
		JOIN recipe_food rf ON mr.recipe_id = rf.recipe_id
		JOIN recipe r ON r.id = mr.recipe_id
		JOIN food f ON rf.food_id = f.id
		WHERE m.owner_id = ?
		AND m.planned_at BETWEEN ? AND ?
		GROUP BY f.name
	`, input.UserID, input.From, input.To)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	if input.Name == "" {
		//input.Name = fmt.Sprintf("Liste: ", userID, input.From, input.To)
		errN := errors.New("shopping list name empty")
		WriteInputError(w, errN)
		return
	}

	var nameArlreadyExist string
	if input.Name != "" {
		_ = s.DB.Get(&nameArlreadyExist, `
			SELECT name
			FROM shopping_list
			WHERE name = ?
			AND user_id = ?
		`, input.Name, userID)
		if nameArlreadyExist != "" {
			errN := errors.New("this list's name already exists")
			WriteInputError(w, errN)
			return
		}
	}

	for _, item := range shoppingList.Items {
		_, err := s.DB.Exec(`
			INSERT INTO shopping_list (food_name, food_quantity, food_unit, fromTime, toTime, name, user_id) VALUES
			(?, ?, ?, ?, ?, ?, ?)
		`, item.Name, item.Quantity, item.Unit, input.From, input.To, input.Name, userID)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
	}

	err = s.DB.Get(&shoppingList, `
		SELECT distinct fromTime, toTime, name, user_id
		FROM shopping_list
		WHERE user_id = ?
		AND fromTime = ?
		AND toTime = ?
		AND name = ?
	`, userID, input.From, input.To, input.Name)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	err = s.DB.Select(&shoppingList.Items, `
		SELECT food_name, food_quantity, food_unit
		FROM shopping_list
		WHERE user_id = ?
		AND fromTime = ?
		AND toTime = ?
	`, userID, input.From, input.To)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	Write(w, shoppingList)
}
