package main

import (
	"net/http"
	"time"
)

type CreateMealInput struct {
	PlannedAt time.Time `json:"planned_at"`
	Guests    int64     `json:"guests"`
	RecipeID  int64     `db:"id" json:"recipe_id"`
}

type CreateMealOutput struct {
	ID int64 `json:"id"`
}

func (s Service) CreateMeal(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input CreateMealInput
	err = Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	res, err := s.DB.Exec(`
		INSERT INTO meal (planned_at, guests, owner_id)
		VALUES (?, ?, ?)
	`, input.PlannedAt, input.Guests, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	mealID, err := res.LastInsertId()
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	_, err = s.DB.Exec(`
		INSERT INTO meal_recipe ( meal_id, recipe_id ) values (?, ?)
	`, mealID, input.RecipeID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	Write(w, CreateMealOutput{ID: mealID})
}
