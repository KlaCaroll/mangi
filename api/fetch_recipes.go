package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/pasztorpisti/qs"
)

type FetchRecipesInput struct {
	Category              []int64 `json:"category,omitempty" qs:"category,omitempty"`
	Name                  string  `json:"name,omitempty" qs:"name,omitempty"`
	UserSettingPreference bool    `json:"preference" qs:"preference"`
	Pagination
}

type FetchRecipesOutput struct {
	Page    int64    `json:"page"`
	Total   int64    `json:"total"`
	Recipes []Recipe `json:"recipes"`
}

func (s Service) FetchRecipes(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != "" {
		s.FetchRecipesAuth(w, r)
	} else {
		s.FetchRecipesAnon(w, r)
	}
}

func (s Service) FetchRecipesAuth(w http.ResponseWriter, r *http.Request) {
	var err error

	userID, err := ParseToken(r, s.secretKey)
	if err != nil {
		WriteUnauthorizedError(w, err)
		return
	}

	var input FetchRecipesInput
	input.PerPage = 10
	err = Read(r, &input)
	if err != nil {
		queryString := fmt.Sprintf(r.URL.RawQuery)
		err = qs.Unmarshal(&input, queryString)
		if err != nil {
			WriteInputError(w, err)
			return
		}

	}

	// Build the conditionnal part of the query, which can then be re-used
	// in the totalQuery and rowQuery.
	// TODO Check the left join for recipes without Category.

	whereClause, a := s.BuildWhereClause(r, input, userID)
	if len(a) != 0 {
		whereClause = " " + whereClause
	}

	var total int64
	CountQuery := fmt.Sprintf(`
		SELECT COUNT(*) 
		FROM recipe r
		JOIN recipe_category rt on r.id = rt.recipe_id 
		WHERE is_public = 1
		%s
		GROUP BY r.id
	`, whereClause)

	CountQuery, args, err := sqlx.In(CountQuery, a...)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}
	err = s.DB.Get(&total, CountQuery, args...)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}
	pages := total / input.PerPage
	if total%input.PerPage != 0 {
		pages += 1
	}
	err = input.Pagination.Validate(pages)
	if err != nil {
		WriteInputError(w, err)
		return
	}

	// OK POUR COUNT

	var recipes []Recipe
	rowsQuery := fmt.Sprintf(`
		SELECT r.id, r.name, r.preparation_time, r.total_time, r.owner_id 
		FROM recipe r
		JOIN recipe_category rt on r.id = rt.recipe_id 
		WHERE is_public = 1
		%s 
		GROUP BY r.id
		LIMIT %v
		OFFSET %v
	`, whereClause, input.Limit(), input.Offset())

	rowsQuery, args, err = sqlx.In(rowsQuery, a...)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}
	err = s.DB.Select(&recipes, rowsQuery, args...)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	Write(w, FetchRecipesOutput{Page: input.Page, Total: pages, Recipes: recipes})

}

func (s Service) FetchRecipesAnon(w http.ResponseWriter, r *http.Request) {
	var err error

	var input FetchRecipesInput
	input.PerPage = 10
	err = Read(r, &input)
	if err != nil {
		queryString := fmt.Sprintf(r.URL.RawQuery)
		err = qs.Unmarshal(&input, queryString)
		if err != nil {
			WriteInputError(w, err)
			return
		}

	}

	if input.Page > 5 {
		WriteUnauthorizedError(w, errors.New("need_to_register_to_see_more"))
	}
	// Build the conditionnal part of the query, which can then be re-used
	// in the totalQuery and rowQuery.
	// TODO Check the left join for recipes without category.

	whereClause, a := s.BuildWhereClauseAnon(r, input)
	if len(a) != 0 {
		whereClause = " " + whereClause
	}

	var total int64
	CountQuery := fmt.Sprintf(`
		SELECT COUNT(*) 
		FROM recipe r
		JOIN recipe_category rt on r.id = rt.recipe_id 
		WHERE is_public = 1
		%s
		GROUP BY r.id
	`, whereClause)

	CountQuery, args, err := sqlx.In(CountQuery, a...)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}
	err = s.DB.Get(&total, CountQuery, args...)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}
	pages := total / input.PerPage
	if total%input.PerPage != 0 {
		pages += 1
	}
	err = input.Pagination.Validate(pages)
	if err != nil {
		WriteInputError(w, err)
		return
	}
	// OK POUR COUNT

	var recipes []Recipe
	rowsQuery := fmt.Sprintf(`
		SELECT r.id, r.name, r.preparation_time, r.total_time, r.owner_id 
		FROM recipe r
		JOIN recipe_category rt on r.id = rt.recipe_id 
		WHERE is_public = 1
		%s 
		GROUP BY r.id
		LIMIT %v
		OFFSET %v
	`, whereClause, input.Limit(), input.Offset())
	rowsQuery, args, err = sqlx.In(rowsQuery, a...)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}
	err = s.DB.Select(&recipes, rowsQuery, args...)
	if err != nil {
		WriteDatabaseError(w, err)
		return
	}

	Write(w, FetchRecipesOutput{Page: input.Page, Total: pages, Recipes: recipes})
}

func (s Service) BuildWhereClause(r *http.Request, input FetchRecipesInput, userID int64) (string, []any) {
	var query string
	var args []any

	if input.Name != "" {
		query += `
			AND r.name LIKE '%` + input.Name + `%'
		`
	}
	if len(input.Category) != 0 {
		query += `
			AND
		`
		for i, reqCategory := range input.Category {
			if i != 0 {
				query += `
					OR
				`
			}
			query += `
				rt.category_id IN (?)
			`
			args = append(args, reqCategory)
		}
	}
	if input.UserSettingPreference {
		// TODO Very inefficient, should be optimized later.
		query += `
			AND NOT r.id IN (
				SELECT r.id
				FROM recipe r
				JOIN recipe_food rf ON rf.recipe_id = r.id
				JOIN food f ON f.id = rf.food_id
				JOIN food_preference fe ON fe.food_id = f.id
				JOIN preference e ON fe.preference_id = e.id
				JOIN user_preference ue ON  ue.preference_id = e.id
				JOIN user us ON ue.user_id = us.id
				WHERE us.id = ?
				AND ue.exist = 1
			)
		`
		args = append(args, userID)
	}
	return query, args
}

func (s Service) BuildWhereClauseAnon(r *http.Request, input FetchRecipesInput) (string, []any) {
	var query string
	var args []any

	if input.Name != "" {
		query += `
			AND r.name LIKE '%` + input.Name + `%'
		`
	}
	if len(input.Category) != 0 {
		query += `
			AND
		`
		for i, reqCategory := range input.Category {
			if i != 0 {
				query += `
					OR
				`
			}
			query += `
				rt.category_id IN (?)
			`
			args = append(args, reqCategory)
		}
	}
	return query, args
}
