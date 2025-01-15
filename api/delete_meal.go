package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/pasztorpisti/qs"
)

type DeleteMealInput struct {
	ID int64 `json:"id" qs:"id"`
}

func (s Service) DeleteMeal(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input DeleteMealInput
	err = Read(r, &input)
	if err != nil {
		queryString := fmt.Sprintf(r.URL.RawQuery)
		err = qs.Unmarshal(&input, queryString)
		if err != nil {
			WriteInputError(w, err)
			return
		}

	}

	var ownerID int64
	err = s.DB.Get(&ownerID, `
		SELECT owner_id
		FROM meal
		WHERE id = ?
	`, input.ID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	// TODO gerer le cas admin

	if ownerID != userID {
		WriteUnauthorizedError(w, errors.New("now the meal owner"))
		return
	}

	_, err = s.DB.Exec(`
		DELETE FROM meal_recipe
		WHERE meal_recipe.meal_id = ?
	`, input.ID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	_, err = s.DB.Exec(`
		DELETE FROM meal
		WHERE id = ?
	`, input.ID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	WriteAck(w)
}
