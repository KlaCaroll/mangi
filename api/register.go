package main

import (
	"errors"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

type RegisterInput struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	Name     string `db:"name" json:"name"`
}

type RegisterOutput struct {
	ID int64 `json:"id"`
}

func (s Service) Register(w http.ResponseWriter, r *http.Request) {
	var input RegisterInput
	err := Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	err = SendMailRegister(input.Email)
	if err != nil {
		WriteInternalErrorMail(w, err)
		return
	}

	passwordHash, err := HashPassword(input.Password)
	if err != nil {
		WritePasswordError(w, err)
		return
	}

	res, err := s.DB.Exec(`
		INSERT INTO user (name, email, password)
		VALUES (?, ?, ?)
	`, input.Name, input.Email, passwordHash)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			WriteError(w, "email_already_exists_error", errors.New("an user with this email already exists"))
			return
		}

		WriteDatabaseError(w, err)
		return
	}

	userID, err := res.LastInsertId()
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	// Seed the preferences by default before user handle his settings
	_, err = s.DB.Exec(`
		INSERT INTO user_ustensil (ustensil_id, user_id) VALUES 
			(1, ?),
			(2, ?),
			(3, ?),
			(4, ?),
			(5, ?),
			(6, ?)
	`, userID, userID, userID, userID, userID, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}
	// Seed the ustensils by default before user handle his settings
	_, err = s.DB.Exec(`
		INSERT INTO user_preference (preference_id, user_id) VALUES 
			(1, ?), (2, ?),
			(3, ?), (4, ?),
			(5, ?), (6, ?),
			(7, ?), (8, ?),
			(9, ?), (10, ?),
			(11, ?), (12, ?),
			(13, ?), (14, ?),
			(15, ?), (16, ?),
			(17, ?), (18, ?),
			(19, ?), (20, ?)	
	`, userID, userID, userID, userID, userID, userID, userID, userID, userID, userID, userID, userID, userID, userID, userID, userID, userID, userID, userID, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	WriteAck(w)
}
