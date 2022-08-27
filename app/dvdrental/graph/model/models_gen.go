// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Actor struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	LastUpdates time.Time `json:"last_updates"`
}

type Customer struct {
	ID         string    `json:"id"`
	StoreID    string    `json:"store_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	AddressID  string    `json:"address_id"`
	CreateDate time.Time `json:"create_date"`
	LastUpdate time.Time `json:"last_update"`
	Active     int       `json:"active"`
}

type Film struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Description     *string   `json:"description"`
	ReleaseYear     int       `json:"release_year"`
	LanguageID      string    `json:"language_id"`
	RentalDuration  int       `json:"rental_duration"`
	RentalRate      string    `json:"rental_rate"`
	Length          *int      `json:"length"`
	ReplacementCost string    `json:"replacement_cost"`
	Rating          string    `json:"rating"`
	LastUpdate      time.Time `json:"last_update"`
	SpecialFeatures []string  `json:"special_features"`
	Fulltext        string    `json:"fulltext"`
}

type Pagination struct {
	Offset *int `json:"offset"`
	Limit  *int `json:"limit"`
}
