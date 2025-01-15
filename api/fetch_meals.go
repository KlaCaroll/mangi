package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pasztorpisti/qs"
)

type FetchMealsInput struct {
	From time.Time `json:"from" qs:"from"`
	To   time.Time `json:"to" sq:"to"`
}

type FetchMealsOutput struct {
	Meals []Meal `json:"meals"`
}

func (s Service) FetchMeals(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input FetchMealsInput
	err = Read(r, &input)
	if err != nil {
		queryString := fmt.Sprintf(r.URL.RawQuery)
		err = qs.Unmarshal(&input, queryString)
		if err != nil {
			WriteInputError(w, err)
			return
		}

	}

	var meals []Meal
	err = s.DB.Select(&meals, `
		SELECT m.id, m.planned_at, m.guests
		FROM meal m
		JOIN user u ON u.id = m.owner_id
		JOIN meal_recipe mr ON mr.meal_id = m.id
		JOIN recipe r ON r.id = mr.recipe_id
		WHERE planned_at BETWEEN ? AND ?
		AND m.owner_id = ?
	`, input.From, input.To, userID)
	if err != nil {
		if len(meals) == 0 {
			WriteError(w, "database_return:", errors.New("no result"))
		} else {
			WriteDatabaseError(w, err)
		}
		return
	}

	if len(meals) == 0 {
		Write(w, meals)
		return
	}

	// TODO (caroll) Revoir la deuxième requête.

	var mealIDs = make([]int64, 0, len(meals))
	var mealsByID = make(map[int64]Meal, len(meals))
	for _, meal := range meals {
		mealIDs = append(mealIDs, meal.ID)
		mealsByID[meal.ID] = meal
	}

	query, args, err := sqlx.In(`
		SELECT mr.meal_id, id, name
		FROM recipe as r
		JOIN meal_recipe as mr ON r.id = mr.recipe_id
		WHERE mr.meal_id in (?)
	`, mealIDs)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	var res []struct {
		MealID int64 `db:"meal_id"`
		Recipe
	}
	err = s.DB.Select(&res, query, args...)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	for _, r := range res {
		meal, ok := mealsByID[r.MealID]
		if !ok {
			WriteDatabaseError(w, err)
			return
		}
		meal.Recipe = append(meal.Recipe, r.Recipe)
		mealsByID[r.MealID] = meal
	}

	meals = make([]Meal, 0, len(meals))
	for _, m := range mealsByID {
		meals = append(meals, m)
	}

	Write(w, FetchMealsOutput{Meals: meals})
}
