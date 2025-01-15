package main

import (
	"fmt"
	"net/http"

	"github.com/pasztorpisti/qs"
)

type homeID struct {
	ID int64 `db:"home_id" qs:"home_id"`
}

func (s Service) FetchHomeShoppingLists(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input homeID
	err = Read(r, &input)
	if err != nil {
		queryString := fmt.Sprintf(r.URL.RawQuery)
		err = qs.Unmarshal(&input, queryString)
		if err != nil {
			WriteInputError(w, err)
			return
		}
	}

	err = s.CheckHomePermission(input.ID, userID)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var lists []UpdateShoppingListInput
	err = s.DB.Select(&lists, `
		SELECT distinct user_id, name, fromTime, toTime, home_id
		FROM shopping_list
		WHERE home_id = ?
	`, input.ID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	Write(w, lists)
}
