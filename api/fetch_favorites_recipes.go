package main

import "net/http"

func (s Service) FetchRecipesFavorites(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var recipes []Recipe
	err = s.DB.Select(&recipes, `
		SELECT id, name, preparation_time, total_time, owner_id
		FROM recipe
		JOIN user_recipe_favorite urf ON urf.recipe_id = recipe.id
		WHERE urf.user_id = ?
	`, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	Write(w, recipes)
}
