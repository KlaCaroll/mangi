package main

import (
	"errors"
	"net/http"
)

func (s Service) FetchShoppingLists(w http.ResponseWriter, r *http.Request) {
	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var homeID []int64
	err = s.DB.Select(&homeID, `
		SELECT home_id
		FROM user_home
		WHERE user_id = ?
	`, userID)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	var ShoppingLists []ShoppingList
	if len(homeID) > 0 {
		for i := range homeID {
			err = s.DB.Select(&ShoppingLists, `
				SELECT user_id, name, fromTime, toTime, home_id
				FROM shopping_list
				WHERE user_id = ?
				UNION 
				SELECT user_id, name, fromTime, toTime, home_id
				FROM shopping_list
				WHERE home_id = ?
			`, &userID, homeID[i])
			if err != nil {
				if len(ShoppingLists) == 0 {
					WriteError(w, "database_return:", errors.New("no result"))
					return
				} else {
					WriteDatabaseError(w, err)
					return
				}
			}
			for i := range ShoppingLists {
				err = s.DB.Select(&ShoppingLists[i].Items, `
					SELECT food_name, food_quantity, food_unit
					FROM shopping_list
					WHERE fromTime = ?
					AND toTime = ?
					AND user_id = ?
				`, ShoppingLists[i].From, ShoppingLists[i].To, ShoppingLists[i].UserID)
				if err != nil {
					WriteDatabaseError(w, err)
					return
				}
			}
			Write(w, ShoppingLists)
			return
		}
	} else if len(homeID) == 0 {
		err = s.DB.Select(&ShoppingLists, `
			SELECT distinct user_id, name, fromTime, toTime, home_id
			FROM shopping_list
			WHERE user_id = ?
		`, userID)
		if err != nil {
			if len(ShoppingLists) == 0 {
				WriteError(w, "database_return:", errors.New("no result"))
				return
			} else {
				WriteDatabaseError(w, err)
				return
			}
		}
		for i := range ShoppingLists {
			err = s.DB.Select(&ShoppingLists[i].Items, `
				SELECT food_name, food_quantity, food_unit
				FROM shopping_list
				WHERE fromTime = ?
				AND toTime = ?
				AND user_id = ?
			`, ShoppingLists[i].From, ShoppingLists[i].To, userID)
			if err != nil {
				WriteDatabaseError(w, err)
				return
			}
		}
		Write(w, ShoppingLists)
		return
	}
}
