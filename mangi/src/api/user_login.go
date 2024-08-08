package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func (s Service) UserLogin(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `db:"email" json:"email"`
		Password string `db:"password" json:"password"`
	}

	err := read(r, &input)
	if err != nil {
		log.Println("parsing input", err)
		writeError(w, "input_error", err)
		return
	}

	var output []struct {
		ID   int64  `db:"id" json:"id"`
		Name string `db:"name" json:"name"`
	}

	err = s.DB.Select(&output, `
		SELECT id, name
		FROM user
		WHERE email = ?
		AND password = ?
	`, input.Email, input.Password)
	if err != nil {
		log.Println("querying database", err)
		writeError(w, "database_error", err)
		return
	}

	write(w, output)
}
