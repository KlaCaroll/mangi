package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/pasztorpisti/qs"
)

type ShowRecipeInput struct {
	ID     int64   `json:"id" qs:"id"`
	Guests float64 `json:"guests" qs:"guests"`
}

type ShowRecipeOutput struct {
	Recipe Recipe `json:"recipe"`
}

func (s Service) ShowRecipe(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input ShowRecipeInput
	err = Read(r, &input)
	if err != nil {
		queryString := fmt.Sprintf(r.URL.RawQuery)
		err = qs.Unmarshal(&input, queryString)
		if err != nil {
			WriteInputError(w, err)
			return
		}

	}

	var recipe Recipe
	err = s.DB.Get(&recipe, `
		SELECT id, name, preparation_time, total_time, description, owner_id, is_public
		FROM recipe 
		WHERE id = ?
	`, input.ID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	if recipe.IsPublic == 0 {
		if recipe.OwnerID != userID {
			errN := errors.New("not_public_recipe")
			WriteUnauthorizedError(w, errN)
			return
		}
	}

	err = s.DB.Get(&recipe.OwnerName, `
		SELECT name
		FROM user 
		WHERE id = ?
	`, recipe.OwnerID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	err = s.DB.Select(&recipe.Ingredients, `
		SELECT f.id, f.name, rf.quantity*? as quantity, rf.unit
		FROM recipe r
		JOIN recipe_food rf ON rf.recipe_id = r.id
		JOIN food f ON f.id = rf.food_id
		WHERE r.id = ?
	`, input.Guests, input.ID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	for i := range recipe.Ingredients {
		recipe.Ingredients[i].Quantity = float64(int(recipe.Ingredients[i].Quantity*100)) / 100
	}

	Write(w, recipe)
}
