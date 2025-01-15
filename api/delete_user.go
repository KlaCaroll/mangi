package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/pasztorpisti/qs"
)

type DeleteUserInput struct {
	ID int64 `json:"id" qs:"id"`
}

func (s Service) UserDelete(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input DeleteUserInput
	err = Read(r, &input)
	if err != nil {
		queryString := fmt.Sprintf(r.URL.RawQuery)
		err = qs.Unmarshal(&input, queryString)
		if err != nil {
			WriteInputError(w, err)
			return
		}
	}
	// gérer le cas admin ET possibilité de récuperer les datas pour l'utilisateur
	if input.ID != userID {
		WriteUnauthorizedError(w, errors.New("can't delete this user"))
		return
	}
	// Check data to delete as recipe that the creator put on is_public = true
	s.checkDataToDelete(userID)

	_, err = s.DB.Exec(`
		DELETE FROM user
		WHERE id = ?
	`, input.ID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	WriteAck(w)
}
