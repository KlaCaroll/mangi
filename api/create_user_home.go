package main

import (
	"errors"
	"fmt"
	"net/http"
)

type CreateHomeInput struct {
	Name string `db:"name" json:"home_name"`
}

func (s Service) CreateUserHome(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input CreateHomeInput
	err = Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	// Check that the home does'nt exist yet.
	// we want to contrain user not to have homes with the same name for 2 reasons.
	// First because interface will be more easy for user if homes have different names
	// and it will be easier to handle the db for the next steps
	var CheckHome Home
	err = s.DB.Get(&CheckHome.ID, `
		SELECT id
		FROM home
		WHERE owner_id = ?
		AND name = ?
	`, userID, input.Name)
	if err != nil {
		CheckHome.ID = 0
	}
	if CheckHome.ID != 0 {
		WriteInputError(w, errors.New("Home already exists"))
		return
	}

	// Now that we now this home doesn't exist for this user we can create a new home.
	res, err := s.DB.Exec(`
		INSERT INTO home (owner_id, name) VALUES (?, ?)
	`, userID, input.Name)
	if err != nil {
		fmt.Println("insert home")
		WriteDatabaseError(w, err)
		return
	}

	var home Home
	home.ID, err = res.LastInsertId()
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	_, err = s.DB.Exec(`
		INSERT INTO user_home (home_id, user_id) VALUES (?, ?)
	`, home.ID, userID)
	if err != nil {
		fmt.Println("checkhome user")
		WriteDatabaseError(w, err)
		return
	}

	err = s.DB.Get(&home, `
		SELECT u.name, h.owner_id
		FROM home h
		JOIN user_home uh ON uh.home_id = h.id
		JOIN user u ON u.id = uh.user_id
		WHERE h.id = ?
	`, home.ID)
	if err != nil {
		fmt.Println("select home")
		WriteDatabaseError(w, err)
		return
	}

	err = s.DB.Get(&home.OwnerName, `
		SELECT name
		FROM user
		WHERE id = ?
	`, home.OwnerID)
	if err != nil {
		fmt.Println("select owner name")
		WriteDatabaseError(w, err)
		return
	}

	Write(w, home)
}
