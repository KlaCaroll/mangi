package main

import "net/http"

type CategoryRecipeOutput struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (s Service) FetchCategoriesList(w http.ResponseWriter, r *http.Request) {
	_, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var categories []CategoryRecipeOutput
	err = s.DB.Select(&categories, `
		SELECT id, name 
		FROM category
	`)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	Write(w, categories)
}
