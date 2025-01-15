package main

import (
	"net/http"
	"time"
)

type DeleteFromShoppingListInput struct {
	From  time.Time `db:"fromTime" json:"from"`
	To    time.Time `db:"toTime" json:"to"`
	Name  string    `db:"name" json:"name,omitempty"`
	Items []struct {
		Name string `db:"food_name" json:"name"`
	} `json:"items"`
}

func (s Service) ShoppingListDeletingItem(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input DeleteFromShoppingListInput
	err = Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	var itemToDelete []int64
	for i := range input.Items {
		var idToPush int64
		err := s.DB.Get(&idToPush, `
			SELECT id
			from shopping_list
			WHERE food_name = ?
			AND fromTime = ?
			AND toTime = ?
			AND user_id = ?
			AND name = ?
		`, input.Items[i].Name, input.From, input.To, userID, input.Name)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
		itemToDelete = append(itemToDelete, idToPush)
	}

	if len(itemToDelete) == 0 {
		WriteDatabaseError(w, err)
		return
	}

	for _, itemID := range itemToDelete {
		err := s.DB.Select(&itemToDelete, `
			DELETE FROM shopping_list
			WHERE id = ?
		`, itemID)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
	}

	var shoppingList ShoppingList
	err = s.DB.Select(&shoppingList.Items, `
		SELECT food_name, food_quantity, food_unit
		FROM shopping_list
		WHERE user_id = ?
		AND fromTime =  ? 
		AND toTime = ?
		AND name = ?
	`, userID, input.From, input.To, input.Name)
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
		AND name = ?
	`, userID, input.From, input.To, input.Name)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	Write(w, shoppingList)
}
