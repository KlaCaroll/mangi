package main

import (
	"fmt"
	"net/http"

	"github.com/pasztorpisti/qs"
)

type ShowUserInput struct {
	ID int64 `json:"id" qs:"id"`
}

type ShowUserOutput struct {
	ID          int64        `db:"id" json:"id"`
	Name        string       `db:"name" json:"name"`
	Email       string       `db:"email" json:"email"`
	Ustensils   []Ustensil   `json:"ustensils"`
	Preferences []Preference `json:"preferences"`
}

func (s Service) ShowUser(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input ShowUserInput
	err = Read(r, &input)
	if err != nil {
		queryString := fmt.Sprintf(r.URL.RawQuery)
		err = qs.Unmarshal(&input, queryString)
		if err != nil {
			WriteInputError(w, err)
			return
		}
	}

	var user ShowUserOutput
	err = s.DB.Get(&user, `
		SELECT id, name, email
		FROM user
		WHERE id = ?
	`, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	err = s.DB.Select(&user.Ustensils, `
		SELECT u.id, u.name
		FROM ustensil u 
		JOIN user_ustensil uu ON uu.ustensil_id = u.id
		WHERE uu.user_id = ?
		AND uu.exist = 1
		GROUP BY u.id
	`, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	err = s.DB.Select(&user.Preferences, `
		SELECT p.id, p.name
		FROM preference p 
		JOIN user_preference up ON up.preference_id = p.id
		WHERE up.user_id = ?
		AND up.exist = 1
		GROUP BY p.id
	`, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	Write(w, user)
}
