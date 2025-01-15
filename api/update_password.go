package main

import (
	"errors"
	"fmt"
	"net/http"
)

type FetchPasswordInput struct {
	OldPassword  string `json:"old_password"`
	NewPassword  string `json:"new_password"`
	Confirmation string `json:"confirmation"`
}

func (s Service) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input FetchPasswordInput
	err = Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	if input.NewPassword != input.Confirmation {
		WriteInputError(w, errors.New("input new passwords mismatch"))
		return
	}

	var DBPassword string
	err = s.DB.Get(&DBPassword, `
		SELECT password
		FROM user
		WHERE id = ?
	`, userID)
	if err != nil {
		if DBPassword == "" {
			WriteError(w, "input_error:", errors.New("wrong password"))
		} else {
			WriteDatabaseError(w, err)
		}
		return
	}

	if !ValidatePassword(input.OldPassword, DBPassword) {
		errN := errors.New("this password doesn't match")
		WriteError(w, "checking_password:", errN)
		fmt.Printf("checking password: %+v\n", err)
		return
	}

	// TODO set a protocol of sendmail type
	var mail string
	err = s.DB.Get(&mail, `
		SELECT email
		FROM user
		WHERE id = ?
	`, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	err = SendMailPassword(mail)
	if err != nil {
		WriteInternalErrorMail(w, err)
		return
	}

	newPasswordHash, err := HashPassword(input.Confirmation)
	if err != nil {
		WritePasswordError(w, err)
		return
	}

	_, err = s.DB.Exec(`
		UPDATE user SET password = ? WHERE id = ? ;
	`, newPasswordHash, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	// TODO send mail with the smtp to the new mail to say password change
	WriteAck(w)
}
