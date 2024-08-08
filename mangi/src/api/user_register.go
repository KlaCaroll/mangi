package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func (s Service) UserRegister(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `db:"email" json:"email"`
		Password string `db:"password" json:"password"`
		Name     string `db:"name" json:"name"`
	}

	err := read(r, &input)
	if err != nil {
		log.Println("parsing input", err)
		writeError(w, "input_error", err)
		return
	}

	res, err := s.DB.Exec(`
		INSERT INTO user (name, email, password)
		VALUES (?, ?, ?)
	`, input.Name, input.Email, input.Password)
	if err != nil {
		log.Println("querying database", err)
		writeError(w, "database_error", err)
		return
	}

	var output struct {
		ID int64 `db:"id" json:"id"`
	}

	output.ID, err = res.LastInsertId()
	if err != nil {
		log.Println("querying database", err)
		writeError(w, "database_error", err)
		return
	}

	write(w, output)
}
