package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pasztorpisti/qs"
)

type ShoppingListInput struct {
	From time.Time `json:"from,omitempty" qs:"from,omitempty"`
	To   time.Time `json:"to,omitempty" qs:"to,omitempty"`
	Name string    `json:"name,omitempty" qs:"name,omitempty"`
}

func (s Service) ShowShoppingList(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input ShoppingListInput
	err = Read(r, &input)
	if err != nil {
		queryString := fmt.Sprintf(r.URL.RawQuery)
		err = qs.Unmarshal(&input, queryString)
		if err != nil {
			WriteInputError(w, err)
			return
		}

	}

	var shoppingList ShoppingList
	if input.Name != "" {
		err = s.DB.Get(&shoppingList, `
			SELECT user_id, name, fromTime, toTime, home_id
			FROM shopping_list
			WHERE name = ?
		`, input.Name)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
	} else {
		err = s.DB.Get(&shoppingList, `
			SELECT user_id, name, fromTime, toTime, home_id
			FROM shopping_list
			WHERE fromTime = ?
			AND toTime = ?
		`, input.From, input.To)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
	}

	if userID != shoppingList.UserID {
		err = s.CheckHomePermission(shoppingList.HomeID, userID)
		if err != nil {
			WriteUnauthorizedError(w, err)
			return
		}
	}

	err = s.DB.Select(&shoppingList.Items, `
		SELECT food_name, food_quantity, food_unit
		FROM shopping_list
		WHERE fromTime = ?
		AND toTime = ?
	`, shoppingList.From, shoppingList.To)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	Write(w, shoppingList)
}
