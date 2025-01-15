package main

import "net/http"

type ItemsOutput struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (s Service) FetchItems(w http.ResponseWriter, r *http.Request) {
	_, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var items []ItemsOutput
	err = s.DB.Select(&items, `
		SELECT id, name 
		FROM food
	`)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	Write(w, items)
}
