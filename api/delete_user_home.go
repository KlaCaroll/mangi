package main

import (
	"net/http"
)

type DeleteHomeInput struct {
	ID       int64  `db:"id" json:"home_id,omitempty"`
	HomeName string `db:"name" json:"home_name,omitempty"`
}

func (s Service) DeleteUserHome(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input DeleteHomeInput
	err = Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	var home Home
	if input.HomeName != "" {
		err = s.DB.Get(&home.ID, `
			SELECT id
			FROM home
			WHERE name = ?
			AND owner_id = ?
		`, input.HomeName, userID)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
	} else {
		err = s.DB.Get(&home.ID, `
			SELECT id
			FROM home
			WHERE id = ?
			AND owner_id = ?
		`, input.ID, userID)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
	}

	if home.ID != 0 {
		_, err = s.DB.Exec(`
			DELETE FROM home WHERE id = ?
		`, home.ID)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
	}

	WriteAck(w)
}
