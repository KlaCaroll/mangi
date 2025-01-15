package main

import (
	"errors"
	"fmt"
	"net/http"
)

type UserUpdatePreferenceInput struct {
	UserID          int64 `db:"id" json:"user_id"`
	VegetarianExist bool  `db:"exist" json:"vegetarian_exist"`
	VeganExist      bool  `db:"exist" json:"Vegan_exist"`
	PorkExist       bool  `db:"exist" json:"Pork_exist"`
	PimentExist     bool  `db:"exist" json:"Piment_exist"`
	BeefExist       bool  `db:"exist" json:"Beef_exist"`
	LactoseExist    bool  `db:"exist" json:"Lactose_exist"`
	PeanutsExist    bool  `db:"exist" json:"Peanuts_exist"`
	GlutenExist     bool  `db:"exist" json:"Gluten_exist"`
	CrustaceanExist bool  `db:"exist" json:"Crustacean_exist"`
	EggExist        bool  `db:"exist" json:"Egg_exist"`
	NutsExist       bool  `db:"exist" json:"Nuts_exist"`
	FructoseExist   bool  `db:"exist" json:"Fructose_exist"`
	SeafoodExist    bool  `db:"exist" json:"Seafood_exist"`
	CeleryExist     bool  `db:"exist" json:"Celery_exist"`
	FishExist       bool  `db:"exist" json:"Fish_exist"`
	MustardExist    bool  `db:"exist" json:"Mustard_exist"`
	SesameExist     bool  `db:"exist" json:"Sesame_exist"`
	SoyExist        bool  `db:"exist" json:"Soy_exist"`
	SulphitesExist  bool  `db:"exist" json:"Sulphites_exist"`
	LupineExist     bool  `db:"exist" json:"Lupine_exist"`
}

func (s Service) UpdateUserPreference(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input UserUpdatePreferenceInput
	err = Read(r, &input)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	if userID != input.UserID {
		WriteUnauthorizedError(w, errors.New("not your profil"))
		return
	}

	err = s.UpdateSetPreference(w, input, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	WriteAck(w)
}

func (s Service) UpdateSetPreference(w http.ResponseWriter, input UserUpdatePreferenceInput, userID int64) error {
	_, err := s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.VegetarianExist, userID, 1)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.VeganExist, userID, 2)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.PorkExist, userID, 3)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.PimentExist, userID, 4)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.BeefExist, userID, 5)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.LactoseExist, userID, 6)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.PeanutsExist, userID, 7)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.GlutenExist, userID, 8)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.CrustaceanExist, userID, 9)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.EggExist, userID, 10)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.NutsExist, userID, 11)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.FructoseExist, userID, 12)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.SeafoodExist, userID, 13)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.CeleryExist, userID, 14)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.FishExist, userID, 15)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.MustardExist, userID, 16)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.SesameExist, userID, 17)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.SoyExist, userID, 18)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.SulphitesExist, userID, 19)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}
	_, err = s.DB.Exec(`
		UPDATE user_preference
		SET 
		 exist = ?
		WHERE user_id = ?
		AND preference_id = ?
	`, input.LupineExist, userID, 20)
	if err != nil {
		Write(w, errors.New("error to update microwave"))
		fmt.Printf("database_error:%+v\n", err)
		return err
	}

	return nil
}
