package main

import (
	"errors"
	"net/http"
	"time"
)

type UpdateMealInput struct {
	MealID      int64     `json:"meal_id"`
	PlannedAt   time.Time `json:"planned_at"`
	Guests      int64     `json:"guests"`
	OldRecipeID int64     `json:"old_recipe_id"`
	RecipeID    int64     `json:"recipe_id"`
}

func (s Service) UpdateMeal(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input UpdateMealInput
	err = Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	var ownerID int64
	err = s.DB.Get(&ownerID, `
		SELECT owner_id
		FROM meal
		WHERE id = ?
	`, input.MealID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	if ownerID != userID {
		errN := errors.New("cant_update_this_recipe")
		WriteError(w, "Unauthoried_request_error", errN)
		return
	}

	_, err = s.DB.Exec(`
		UPDATE meal
		SET 
		 planned_at = ?,
		 guests = ?
		WHERE id = ?
	`, input.PlannedAt, input.Guests, input.MealID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	// avoir la jonction pour la changer
	_, err = s.DB.Exec(`
		UPDATE meal_recipe
		SET
		 recipe_id = ?
		WHERE meal_recipe.recipe_id = ? AND meal_recipe.meal_id = ?
	`, input.RecipeID, input.OldRecipeID, input.MealID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	WriteAck(w)
}
