package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type UpdateShoppingListInput struct {
	UserID int64              `db:"user_id" json:"user_id"`
	From   time.Time          `db:"fromTime" json:"from"`
	To     time.Time          `db:"toTime" json:"to"`
	Name   string             `db:"name" json:"name,omitempty"`
	HomeID int64              `db:"home_id" json:"home_id,omitempty"`
	Items  []ShoppingListItem `json:"items"`
}

type UpdateShoppingListOutput struct {
	UserID int64              `db:"user_id" json:"user_id"`
	From   time.Time          `db:"fromTime" json:"from"`
	To     time.Time          `db:"toTime" json:"to"`
	Name   string             `db:"name" json:"name"`
	HomeID int64              `db:"home_id" json:"home_id,omitempty"`
	Items  []ShoppingListItem `json:"items"`
}

func (s Service) UpdateShoppingList(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input UpdateShoppingListInput
	err = Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	err = s.DeleteList(r, input, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	shoppingList, err := s.CreateList(r, input)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	Write(w, shoppingList)

}

func (s Service) CreateList(r *http.Request, input UpdateShoppingListInput) (UpdateShoppingListOutput, error) {
	var shoppingList UpdateShoppingListOutput
	if input.Name == "" {
		input.Name = fmt.Sprintf("%v-%s-to-%s", input.UserID, input.From, input.To)
	}
	for _, item := range input.Items {
		_, err := s.DB.Exec(`
			INSERT INTO shopping_list (food_name, food_quantity, food_unit, fromTime, toTime, name, user_id, home_id) VALUES
			(?, ?, ?, ?, ?, ?, ?, ?)
		`, item.Name, item.Quantity, item.Unit, input.From, input.To, input.Name, input.UserID, input.HomeID)
		if err != nil {
			fmt.Printf("database_error_insert: %+v", err)
			return shoppingList, errors.New("database_error")
		}
	}

	err := s.DB.Get(&shoppingList, `
		SELECT fromTime, toTime, name, user_id, home_id
		FROM shopping_list
		WHERE user_id = ?
		AND fromTime = ?
		AND toTime = ?
	`, input.UserID, input.From, input.To)
	if err != nil {
		fmt.Printf("database_error_select: %+v", err)
		return shoppingList, errors.New("database_error")
	}

	err = s.DB.Select(&shoppingList.Items, `
		SELECT food_name, food_quantity, food_unit
		FROM shopping_list
		WHERE user_id = ?
		AND fromTime = ?
		AND toTime = ?
	`, input.UserID, input.From, input.To)
	if err != nil {
		fmt.Printf("database_error_select_item: %+v", err)
		return shoppingList, errors.New("database_error")
	}

	return shoppingList, nil

}

func (s Service) DeleteList(r *http.Request, input UpdateShoppingListInput, userID int64) error {
	var shoppingList ShoppingList
	err := s.DB.Get(&shoppingList, `
		SELECT user_id, fromTime, toTime, home_id
		FROM shopping_list
		WHERE fromTime = ?
		AND toTime = ?
	`, input.From, input.To)
	if err != nil {
		return errors.New("no_shopping_list")
	}

	if userID != shoppingList.UserID {
		err = s.CheckHomePermission(shoppingList.HomeID, userID)
		if err != nil {
			return err
		}
	}

	_, err = s.DB.Exec(`
		DELETE FROM shopping_list
		WHERE fromTime = ?
		AND toTime = ?
	`, shoppingList.From, shoppingList.To)
	if err != nil {
		return err
	}

	return nil
}
