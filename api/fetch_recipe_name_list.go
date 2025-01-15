package main

import "net/http"

type RecipesNamesOutput struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (s Service) FetchRecipeNameList(w http.ResponseWriter, r *http.Request) {
	_, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var lists []RecipesNamesOutput
	err = s.DB.Select(&lists, `
		SELECT id, name 
		FROM recipe
	`)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	Write(w, lists)
}
