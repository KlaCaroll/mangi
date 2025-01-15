package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/pasztorpisti/qs"
)

type DeleteRecipeInput struct {
	ID int64 `json:"id" qs:"id"`
}

func (s Service) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input DeleteRecipeInput
	err = Read(r, &input)
	if err != nil {
		queryString := fmt.Sprintf(r.URL.RawQuery)
		err = qs.Unmarshal(&input, queryString)
		if err != nil {
			WriteInputError(w, err)
			return
		}
	}

	// TODO (caroll) Ajouter le user_id aux recipes pour permettre de vérifier les permissions.
	var ownerID []int64
	err = s.DB.Select(&ownerID, `
		SELECT owner_id
		FROM recipe
		WHERE id = ?
	`, input.ID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	if len(ownerID) == 0 {
		WriteError(w, "unauthorized_error", errors.New("not your recipe"))
		return
	}

	if ownerID[0] != userID {
		WriteError(w, "unauthorized_error", errors.New("not your recipe"))
		return
	}

	_, err = s.DB.Exec(`
		DELETE FROM recipe
		WHERE id = ?
	`, input.ID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	WriteAck(w)
}
