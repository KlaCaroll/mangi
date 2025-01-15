package main

import (
	"errors"
	"net/http"
)

type InvitationInput struct {
	ID         int64  `db:"id" json:"home_id,omitempty"`
	Name       string `db:"name" json:"home_name,omitempty"`
	MemberMail string `json:"invitation_to"`
}

func (s Service) InviteUserHome(w http.ResponseWriter, r *http.Request) {
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

	if input.Name == "" && input.ID == 0 {
		WriteInputError(w, errors.New("home information missing"))
		return
	}

	// We have to check that user is authorized to handle this home.
	// we check the junction then if he is the owner member.
	home, err := s.CheckHomeInformations(w, input, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	// All check are passed so we can handle for now the invitation.
	err = SendMailInvitation(input.MemberMail, home.Name, home.OwnerName, home.ID)
	if err != nil {
		WriteInternalErrorMail(w, err)
		return
	}

	err = s.TempAcceptInvitation(w, userID, input.ID, input.MemberMail)
	if err != nil {
		WriteError(w, "accepte_invitation_error:", err)
		return
	}

	Write(w, home)
}

func (s Service) CheckHomeInformations(w http.ResponseWriter, input InvitationInput, userID int64) (Home, error) {
	var home Home
	var err error = nil
	if input.Name != "" {
		err := s.DB.Get(&home, `
			SELECT h.id, h.name, h.owner_id
			FROM home h 
			JOIN user_home uh ON uh.home_id = h.id
			WHERE h.name = ?
			AND uh.user_id = ?
		`, input.Name, userID)
		if err != nil {
			if home.OwnerID != userID {
				WriteUnauthorizedError(w, errors.New("not your home"))
				return home, err
			}
			WriteDatabaseError(w, err)
			return home, err
		}
	} else {
		err := s.DB.Get(&home, `
			SELECT h.id, h.name, h.owner_id
			FROM home h 
			JOIN user_home uh ON uh.home_id = h.id
			JOIN user u ON u.id = uh.user_id
			WHERE h.id = ?
			AND uh.user_id = ?
		`, input.ID, userID)
		if err != nil {
			if home.OwnerID != userID {
				WriteUnauthorizedError(w, errors.New("not your home"))
				return home, err
			}
			WriteDatabaseError(w, err)
			return home, err
		}
		err = s.DB.Get(&home.OwnerName, `
			SELECT name as owner_name
			FROM user
			WHERE id = ?
		`, home.OwnerID)
		if err != nil {
			WriteDatabaseError(w, err)
			return home, err
		}
	}
	return home, err
}

func (s Service) TempAcceptInvitation(w http.ResponseWriter, userID, inputID int64, memberMail string) error {
	var err error = nil

	var InvitedMemberID int64
	err = s.DB.Get(&InvitedMemberID, `
		SELECT id 
		FROM user
		WHERE email = ?
	`, memberMail)
	if err != nil {
		WriteDatabaseError(w, err)
		return err
	}

	_, err = s.DB.Exec(`
		INSERT INTO user_home (user_id, home_id) VALUES (?, ?)
	`, InvitedMemberID, inputID)
	if err != nil {
		WriteDatabaseError(w, err)
		return err
	}

	var home Home
	err = s.DB.Get(&home, `
		SELECT h.id, h.name, h.owner_id
		FROM home h 
		JOIN user_home uh ON uh.home_id = h.id
		JOIN user u ON u.id = uh.user_id
		WHERE h.id = ?
	`, inputID)
	if err != nil {
		WriteDatabaseError(w, err)
		return err
	}
	err = s.DB.Get(&home.OwnerName, `
		SELECT name as owner_name
		FROM user
		WHERE id = ?
	`, home.OwnerID)
	if err != nil {
		WriteDatabaseError(w, err)
		return err
	}

	return err
}
