package main

import "net/http"

type UstensilsOutput struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (s Service) FetchUstensilsList(w http.ResponseWriter, r *http.Request) {
	_, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var ustensils []UstensilsOutput
	err = s.DB.Select(&ustensils, `
		SELECT id, name 
		FROM ustensil
	`)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	Write(w, ustensils)
}
