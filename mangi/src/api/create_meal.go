package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func (s Service) CreateMeal(w http.ResponseWriter, r *http.Request) {
	var input struct {
		PlannedAt time.Time `db:"planned_at" json:"planned_at"`
		Guests    int64     `db:"guests" json:"guests"`
		UserID    int64     `db:"user_id" json:"user_id"`
	}

	err := read(r, &input)
	if err != nil {
		log.Println("parsing input", err)
		writeError(w, "input_error", err)
		return
	}

	res, err := s.DB.Exec(`
		INSERT INTO meal (planned_at, guests, user_id)
		VALUES (?, ?, ?)
	`, input.PlannedAt, input.Guests, input.UserID)
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
