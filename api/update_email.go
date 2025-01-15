package main

import (
	"errors"
	"fmt"
	"net/http"
)

type InputUserInformation struct {
	OldEmail string `db:"email" json:"old_email"`
	NewEmail string `db:"email" json:"new_email"`
}

func (s Service) UpdateEmail(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input InputUserInformation
	err = Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	var oldMailDB string
	err = s.DB.Get(&oldMailDB, `
		SELECT email
		FROM user
		WHERE id = ?
	`, userID)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	if input.OldEmail != oldMailDB {
		errN := errors.New("wrong_input")
		WriteInputError(w, errN)
		return
	}

	// TODO add the smtp service to send mail something happen to new email
	err = SendMailEmail(input.NewEmail)
	if err != nil {
		WriteInternalErrorMail(w, err)
		return
	}

	_, err = s.DB.Exec(`
		UPDATE user SET email = ? WHERE id = ? ;
	`, input.NewEmail, userID)
	if err != nil {
		Write(w, errors.New("email already exists or database error"))
		fmt.Printf("database_error:%+v\n", err)
		return
	}

	// TODO send mail with the smtp to the new mail to say this is the new mail
	WriteAck(w)
}
