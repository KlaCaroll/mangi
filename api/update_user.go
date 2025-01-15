package main

import (
	"net/http"
)

// TODO (caroll) Add a change_password endpoint.
// TODO (caroll) Add a change_email endpoint.

type UserUpdateInput struct {
	Name string `json:"name"`
}

func (s Service) UserUpdate(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input UserUpdateInput
	err = Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	_, err = s.DB.Exec(`
		UPDATE user
		SET 
		 name = ?
		WHERE id = ?
	`, input.Name, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	WriteAck(w)
}
