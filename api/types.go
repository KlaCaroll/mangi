package main

import (
	"errors"
	"time"
)

type Pagination struct {
	Page    int64 `json:"page" qs:"page"`
	PerPage int64
}

func (p Pagination) Validate(total int64) error {
	if p.Page <= 0 {
		return errors.New("invalid page")
	}

	if p.PerPage <= 0 {
		return errors.New("invalid page size")
	}

	if p.Page > total {
		return errors.New("page doesn't exists")
	}

	return nil
}

func (p Pagination) Limit() int64 {
	return p.PerPage
}

func (p Pagination) Offset() int64 {
	return (p.Page - 1) * p.PerPage
}

type ShoppingListItem struct {
	Name     string  `db:"food_name" json:"name"`
	Quantity float64 `db:"food_quantity" json:"quantity"`
	Unit     string  `db:"food_unit" json:"unit"`
}

type ShoppingList struct {
	Name   string             `db:"name" json:"name"`
	UserID int64              `db:"user_id" json:"user_id"`
	From   time.Time          `db:"fromTime" json:"from"`
	To     time.Time          `db:"toTime" json:"to"`
	HomeID any                `db:"home_id" json:"home_id"`
	Items  []ShoppingListItem `json:"items"`
}

type Meal struct {
	ID        int64     `db:"id" json:"id"`
	PlannedAt time.Time `db:"planned_at" json:"planned_at"`
	Guests    uint      `db:"guests" json:"guests"`
	Recipe    []Recipe  `json:"recipes"`
}
type Ustensil struct {
	UstensilID   int64  `db:"id" json:"ustensil_id"`
	UstensilName string `db:"name" json:"ustensil_name"`
	Exist        bool   `db:"exist" json:"exist"`
}

type Category struct {
	CategoryID   int64  `db:"id" json:"category_id"`
	CategoryName string `db:"name" json:"category_name"`
}

type Preference struct {
	PreferenceID   int64  `db:"id" json:"preference_id"`
	PreferenceName string `db:"name" json:"preference_name"`
	Exist          bool   `db:"exist" json:"exist"`
}

type Recipe struct {
	ID              int64        `db:"id" json:"id"`
	Name            string       `db:"name" json:"name"`
	PreparationTime int64        `db:"preparation_time" json:"preparation_time,omitempty"`
	TotalTime       int64        `db:"total_time" json:"total_time,omitempty"`
	Description     string       `db:"description" json:"description"`
	IsPublic        int64        `db:"is_public" json:"is_public"`
	OwnerID         int64        `db:"owner_id" json:"by,omitempty"`
	OwnerName       string       `db:"name" json:"owner,omitempty"`
	Ingredients     []Ingredient `json:"ingredients,omitempty"`
	Ustensils       []Ustensil   `json:"ustensils,omitempty"`
	Categories      []Category   `json:"categories,omitempty"`
}

type Ingredient struct {
	ID       int64   `db:"id" json:"id"`
	Name     string  `db:"name" json:"name"`
	Quantity float64 `db:"quantity" json:"quantity"`
	Unit     string  `db:"unit" json:"unit"`
}

type Member struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Home struct {
	ID        int64    `db:"id" json:"id"`
	Name      string   `db:"name" json:"name"`
	OwnerID   int64    `db:"owner_id" json:"owner_id"`
	OwnerName string   `db:"owner_name" json:"owner_name"`
	Members   []Member `json:"members"`
}

type Comment struct {
	ID          int64  `db:"id" json:"id"`
	OwnerID     int64  `db:"owner_id" json:"owner_id"`
	Description string `db:"description" json:"description"`
	ParentID    int64  `db:"parent_id" json:"parent_id"`
}
