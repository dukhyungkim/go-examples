// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package repository

import (
	"context"

	"github.com/lib/pq"
)

const getActor = `-- name: GetActor :one
SELECT actor_id, first_name, last_name, last_update
FROM actor
WHERE actor_id = $1
LIMIT 1
`

func (q *Queries) GetActor(ctx context.Context, actorID int32) (*Actor, error) {
	row := q.db.QueryRowContext(ctx, getActor, actorID)
	var i Actor
	err := row.Scan(
		&i.ActorID,
		&i.FirstName,
		&i.LastName,
		&i.LastUpdate,
	)
	return &i, err
}

const getCustomer = `-- name: GetCustomer :one
SELECT customer_id, store_id, first_name, last_name, email, address_id, activebool, create_date, last_update, active
FROM customer
WHERE customer_id = $1
LIMIT 1
`

func (q *Queries) GetCustomer(ctx context.Context, customerID int32) (*Customer, error) {
	row := q.db.QueryRowContext(ctx, getCustomer, customerID)
	var i Customer
	err := row.Scan(
		&i.CustomerID,
		&i.StoreID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.AddressID,
		&i.Activebool,
		&i.CreateDate,
		&i.LastUpdate,
		&i.Active,
	)
	return &i, err
}

const getFilm = `-- name: GetFilm :one
SELECT film_id, title, description, release_year, language_id, rental_duration, rental_rate, length, replacement_cost, rating, last_update, special_features, fulltext
FROM film
WHERE film_id = $1
LIMIT 1
`

func (q *Queries) GetFilm(ctx context.Context, filmID int32) (*Film, error) {
	row := q.db.QueryRowContext(ctx, getFilm, filmID)
	var i Film
	err := row.Scan(
		&i.FilmID,
		&i.Title,
		&i.Description,
		&i.ReleaseYear,
		&i.LanguageID,
		&i.RentalDuration,
		&i.RentalRate,
		&i.Length,
		&i.ReplacementCost,
		&i.Rating,
		&i.LastUpdate,
		pq.Array(&i.SpecialFeatures),
		&i.Fulltext,
	)
	return &i, err
}

const listActors = `-- name: ListActors :many
SELECT actor_id, first_name, last_name, last_update
FROM actor
ORDER BY actor_id
LIMIT $1 OFFSET $2
`

type ListActorsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListActors(ctx context.Context, arg *ListActorsParams) ([]*Actor, error) {
	rows, err := q.db.QueryContext(ctx, listActors, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Actor
	for rows.Next() {
		var i Actor
		if err := rows.Scan(
			&i.ActorID,
			&i.FirstName,
			&i.LastName,
			&i.LastUpdate,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCustomers = `-- name: ListCustomers :many
SELECT customer_id, store_id, first_name, last_name, email, address_id, activebool, create_date, last_update, active
FROM customer
ORDER BY customer_id
LIMIT $1 OFFSET $2
`

type ListCustomersParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListCustomers(ctx context.Context, arg *ListCustomersParams) ([]*Customer, error) {
	rows, err := q.db.QueryContext(ctx, listCustomers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Customer
	for rows.Next() {
		var i Customer
		if err := rows.Scan(
			&i.CustomerID,
			&i.StoreID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.AddressID,
			&i.Activebool,
			&i.CreateDate,
			&i.LastUpdate,
			&i.Active,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listFilms = `-- name: ListFilms :many
SELECT film_id, title, description, release_year, language_id, rental_duration, rental_rate, length, replacement_cost, rating, last_update, special_features, fulltext
FROM film
ORDER BY film_id
LIMIT $1 OFFSET $2
`

type ListFilmsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListFilms(ctx context.Context, arg *ListFilmsParams) ([]*Film, error) {
	rows, err := q.db.QueryContext(ctx, listFilms, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Film
	for rows.Next() {
		var i Film
		if err := rows.Scan(
			&i.FilmID,
			&i.Title,
			&i.Description,
			&i.ReleaseYear,
			&i.LanguageID,
			&i.RentalDuration,
			&i.RentalRate,
			&i.Length,
			&i.ReplacementCost,
			&i.Rating,
			&i.LastUpdate,
			pq.Array(&i.SpecialFeatures),
			&i.Fulltext,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
