package main

import (
	"errors"
	"net/http"
)

type UpdateRecipeInput struct {
	ID              int64  `db:"id" json:"id"`
	Name            string `db:"name" json:"name"`
	PreparationTime int64  `db:"preparation_time" json:"preparation_time"`
	TotalTime       int64  `db:"total_time" json:"total_time"`
	Description     string `db:"description" json:"description"`
	IsPublic        int    `db:"is_public" json:"is_public"`
	Ingredients     []struct {
		FoodID   int64   `db:"food_id" json:"food_id"`
		Quantity float64 `db:"quantity" json:"quantity"`
		Unit     string  `db:"unit" json:"unit"`
	} `json:"ingredients"`
}

func (s Service) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	var err error

	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	// TODO (caroll) Après avoir ajouté user_id, vérifier token & owner.
	var input UpdateRecipeInput
	err = Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	var ownerID int64
	err = s.DB.Get(&ownerID, `
		SELECT owner_id
		FROM recipe
		WHERE id = ?
	`, input.ID)
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
		UPDATE recipe
		SET 
		 name = ?, 
		 preparation_time = ?, 
		 total_time = ?, 
		 description = ?,
		 is_public = ?
		WHERE id = ?
		AND owner_id = ?
	`, input.Name, input.PreparationTime, input.TotalTime, input.Description, input.IsPublic, input.ID, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	for _, insert := range input.Ingredients {
		_, err = s.DB.Exec(`
			UPDATE recipe_food
			SET
			 food_id = ?,
			 quantity = ?,
			 unit = ?
			WHERE recipe_id = ?
		`, insert.FoodID, insert.Quantity, insert.Unit, input.ID)
		if err != nil {
			WriteDatabaseError(w, err)
			return
		}
	}

	WriteAck(w)
}
