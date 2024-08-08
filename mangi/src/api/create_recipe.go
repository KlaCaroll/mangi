package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func (s Service) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}

	err := read(r, &input)
	if err != nil {
		log.Println("parsing input", err)
		writeError(w, "input_error", err)
		return
	}

	res, err := s.DB.Exec(`
		INSERT INTO recipe (name)
		VALUES (?)
	`, input.Name)
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
