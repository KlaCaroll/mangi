package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type DeleteShoppingListInput struct {
	UserID int64     `json:"user_id" qs:"user_id"`
	From   time.Time `json:"from" qs:"from"`
	To     time.Time `json:"to" qs:"to"`
	Name   string    `json:"name,omitempty" qs:"name,omitempty"`
}

func (s Service) DeleteShoppingList(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input DeleteShoppingListInput
	err = Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	var shoppingList ShoppingList
	err = s.DB.Get(&shoppingList, `
		SELECT user_id, fromTime, toTime, home_id
		FROM shopping_list
		WHERE fromTime = ?
		AND toTime = ?
	`, input.From, input.To)
	if err != nil {
		WriteError(w, "input_error", errors.New("no_shopping_list"))
		return
	}

	fmt.Printf("user id 1: %+v\n", userID)

	if userID != shoppingList.UserID {
		err = s.CheckHomePermission(shoppingList.HomeID, userID)
		if err != nil {
			WriteUnauthorizedError(w, err)
			return
		}
	}

	fmt.Printf("user id 2: %+v\n", userID)

	_, err = s.DB.Exec(`
		DELETE FROM shopping_list
		WHERE user_id = ?
		AND fromTime = ?
		AND toTime = ?
	`, shoppingList.UserID, input.From, input.To)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	WriteAck(w)
}
