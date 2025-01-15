package main

import "net/http"

type FavoriteRecipeInput struct {
	RecipeID int64 `db:"id" json:"recipe_id"`
}

func (s Service) CreateFavoritesRecipes(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input FavoriteRecipeInput
	err = Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	_, err = s.DB.Exec(`
		INSERT INTO user_recipe_favorite (user_id, recipe_id)
		VALUES (?, ?)
	`, userID, input.RecipeID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	WriteAck(w)
}
