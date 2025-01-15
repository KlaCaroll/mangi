package main

import (
	"net/http"
)

type ShowHomeInput struct {
	ID   int64  `db:"id" json:"home_id,omitempty"`
	Name string `db:"name" json:"home_name,omitempty"`
}

func (s Service) ShowHome(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input InvitationInput
	err = Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	// We have to check that user is authorized to handle this home.
	// we check the junction then if he is the owner or member.
	// information of A home
	home, err := s.CheckHomeInformations(w, input, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}
	err = s.DB.Get(&home.OwnerName, `
		SELECT name as owner_name
		FROM user
		WHERE id = ?
	`, home.OwnerID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	// Get all members of a home.
	err = s.DB.Select(&home.Members, `
		SELECT u.id, u.name
		FROM user u
		JOIN user_home uh ON uh.user_id = u.id
		WHERE uh.home_id = ?
	`, home.ID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	err = s.CheckHomePermission(home.ID, userID)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	Write(w, home)
}
