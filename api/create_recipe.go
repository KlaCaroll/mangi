package main

import (
	"net/http"
)

type CreateRecipeOutput struct {
	ID int64 `json:"id"`
}

func (s Service) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input Recipe
	err = Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	res, err := s.DB.Exec(`
		INSERT INTO recipe (name, preparation_time, total_time, description, owner_id, is_public)
		VALUES (?, ?, ?, ?, ?, ?)
	`, input.Name, input.PreparationTime, input.TotalTime, input.Description, userID, input.IsPublic)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	recipeID, err := res.LastInsertId()
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	for _, categoryR := range input.Categories {
		_, err = s.DB.Exec(`
			INSERT INTO recipe_category (recipe_id, category_id) VALUES (?, ?)
		`, recipeID, categoryR.CategoryID)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
	}

	for _, ustensil := range input.Ustensils {
		_, err = s.DB.Exec(`
			INSERT INTO ustensil_recipe (recipe_id, ustensil_id) VALUES (?, ?)
		`, recipeID, ustensil.UstensilID)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
	}

	for _, insert := range input.Ingredients {
		_, err = s.DB.Exec(`
			INSERT INTO recipe_food (recipe_id, food_id, quantity, unit) 
			VALUES (?, ?, ?, ?)
		`, recipeID, insert.ID, insert.Quantity, insert.Unit)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
	}

	_, err = s.DB.Exec(`
		INSERT INTO user_recipe_favorite (user_id, recipe_id)
		VALUES (?, ?)
	`, userID, recipeID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	Write(w, CreateRecipeOutput{ID: recipeID})
}
