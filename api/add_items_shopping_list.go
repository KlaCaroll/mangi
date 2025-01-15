package main

import (
	"net/http"
	"time"
)

type AddShoppingListInput struct {
	From  time.Time          `db:"fromTime" json:"from"`
	To    time.Time          `db:"toTime" json:"to"`
	Name  string             `db:"name" json:"name,omitempty"`
	Items []ShoppingListItem `json:"items"`
}

func (s Service) ShoppingListAddingItem(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input AddShoppingListInput
	err = Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	var shoppingList ShoppingList
	err = s.DB.Select(&shoppingList.Items, `
		SELECT food_name, food_quantity, food_unit
		FROM shopping_list
		WHERE user_id = ?
		AND fromTime =  ? 
		AND toTime = ?
	`, userID, input.From, input.To)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	err = s.DB.Get(&shoppingList, `
		SELECT fromTime, toTime, name, user_id, home_id
		FROM shopping_list
		WHERE user_id = ?
		AND fromTime = ?
		AND toTime = ?
	`, userID, input.From, input.To)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	err = s.CheckHomePermission(shoppingList.HomeID, userID)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	for _, item := range input.Items {
		_, err := s.DB.Exec(`
			INSERT INTO shopping_list (food_name, food_quantity, food_unit, fromTime, toTime, name, user_id, home_id) VALUES
			(?, ?, ?, ?, ?, ?, ?, ?)
		`, item.Name, item.Quantity, item.Unit, input.From, input.To, shoppingList.Name, userID, shoppingList.HomeID)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
	}

	err = s.DB.Select(&shoppingList.Items, `
		SELECT food_name, SUM(food_quantity) AS food_quantity, MIN(food_unit) AS food_unit
		FROM shopping_list
		WHERE user_id = ?
		AND fromTime =  ? 
		AND toTime = ?
		GROUP BY food_name
	`, userID, input.From, input.To)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	Write(w, shoppingList)
}
