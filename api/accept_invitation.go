package main

import "net/http"

type AcceptInvitationInput struct {
	HomeID int64 `db:"id" json:"home_id"`
}

func (s Service) AcceptHomeInvitation(w http.ResponseWriter, r *http.Request) {
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

	res, err := s.DB.Exec(`
		INSERT INTO user_home (user_id, home_id) VALUES (?, ?)
	`, userID, input.ID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	homeID, err := res.LastInsertId()
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	var home Home
	err = s.DB.Get(&home, `
		SELECT h.id, h.name, h.owner_id
		FROM home h 
		JOIN user_home uh ON uh.home_id = h.id
		JOIN user u ON u.id = uh.user_id
		WHERE h.id = ?
	`, homeID)
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

	Write(w, home)
}
