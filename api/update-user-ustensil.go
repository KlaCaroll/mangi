package main

import (
	"errors"
	"fmt"
	"net/http"
)

type UserUpdateUstensilInput struct {
	UserID              int64 `db:"id" json:"user_id"`
	MicrowaveExist      bool  `db:"exist" json:"microwave_exist"`
	OvenExist           bool  `db:"exist" json:"oven_exist"`
	PressureCookerExist bool  `db:"exist" json:"pressure_cooker_exist"`
	WokExist            bool  `db:"exist" json:"wok_exist"`
	FryerExist          bool  `db:"exist" json:"fryer_exist"`
	MixExist            bool  `db:"exist" json:"mix_exist"`
}

func (s Service) UpdateUserUstensil(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input UserUpdateUstensilInput
	err = Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	if userID != input.UserID {
		WriteUnauthorizedError(w, errors.New("not your profil"))
		return
	}

	err = s.UpdateSetUstensil(w, input, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	WriteAck(w)

}

func (s Service) UpdateSetUstensil(w http.ResponseWriter, input UserUpdateUstensilInput, userID int64) error {
	_, err := s.DB.Exec(`
		UPDATE user_ustensil
		SET 
		 exist = ?
		WHERE user_id = ?
		AND ustensil_id = ?
	`, input.MicrowaveExist, userID, 1)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_ustensil
		SET 
		 exist = ?
		WHERE user_id = ?
		AND ustensil_id = ?
	`, input.OvenExist, userID, 2)
	if err != nil {
		Write(w, errors.New("error to update oven"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_ustensil
		SET 
		 exist = ?
		WHERE user_id = ?
		AND ustensil_id = ?
	`, input.PressureCookerExist, userID, 3)
	if err != nil {
		Write(w, errors.New("error to update pressure_cooker"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_ustensil
		SET 
		 exist = ?
		WHERE user_id = ?
		AND ustensil_id = ?
	`, input.WokExist, userID, 4)
	if err != nil {
		Write(w, errors.New("error to update wok"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_ustensil
		SET 
		 exist = ?
		WHERE user_id = ?
		AND ustensil_id = ?
	`, input.FryerExist, userID, 5)
	if err != nil {
		Write(w, errors.New("error to update fryer"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_ustensil
		SET 
		 exist = ?
		WHERE user_id = ?
		AND ustensil_id = ?
	`, input.MixExist, userID, 6)
	if err != nil {
		Write(w, errors.New("error to update mix"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}

	return nil
}
