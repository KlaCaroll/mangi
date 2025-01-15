package main

import (
	"net/http"
)

func (s Service) FetchHomes(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var homes []Home
	err = s.DB.Select(&homes, `
		SELECT h.id, h.name, h.owner_id, u.name as owner_name
		FROM home h 
		JOIN user_home uh ON uh.home_id = h.id
		JOIN user u ON u.id = uh.user_id
		WHERE uh.user_id = ?
	`, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	var members []Member
	for i, home := range homes {
		err = s.CheckHomePermission(home.ID, userID)
		if err != nil {
			WriteUnauthorizedError(w, err)
			return
		}
		err = s.DB.Select(&members, `
			SELECT name 
			FROM user u
			JOIN user_home uh ON uh.user_id = u.id
			WHERE uh.home_id = ?
		`, home.ID)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
		homes[i].Members = append(homes[i].Members, members...)
	}

	Write(w, homes)
}
