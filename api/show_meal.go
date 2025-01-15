package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/pasztorpisti/qs"
)

type ShowMealInput struct {
	MealID int64 `db:"id" json:"id" qs:"id"`
}

type ShowMealOutput struct {
	MealID     int64     `db:"id" json:"meal_id"`
	PlannedAt  time.Time `db:"planned_at" json:"planned_at"`
	Guests     int64     `db:"guests" json:"guests"`
	RecipeName string    `db:"name" json:"recipe"`
	OwnerID    int64     `db:"owner_id"`
}

func (s Service) ShowMeal(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input ShowMealInput
	err = Read(r, &input)
	if err != nil {
		queryString := fmt.Sprintf(r.URL.RawQuery)
		err = qs.Unmarshal(&input, queryString)
		if err != nil {
			WriteInputError(w, err)
			return
		}

	}

	var meal ShowMealOutput
	err = s.DB.Get(&meal, `
		SELECT r.name, m.planned_at, m.guests, m.id, m.owner_id
		FROM meal m
		JOIN meal_recipe mr ON mr.meal_id = m.id
		JOIN recipe r ON r.id = mr.recipe_id  
		WHERE m.id = ? 
	`, input.MealID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	if userID != meal.OwnerID {
		errN := errors.New("not_your_meal")
		WriteUnauthorizedError(w, errN)
		return
	}

	Write(w, meal)
}
