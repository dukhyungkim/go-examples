-- name: GetActor :one
SELECT *
FROM actor
WHERE actor_id = $1
LIMIT 1;

-- name: ListActors :many
SELECT *
FROM actor
ORDER BY actor_id
LIMIT $1 OFFSET $2;

-- name: GetFilm :one
SELECT *
FROM film
WHERE film_id = $1
LIMIT 1;

-- name: ListFilms :many
SELECT *
FROM film
ORDER BY film_id
LIMIT $1 OFFSET $2;

-- name: GetCustomer :one
SELECT *
FROM customer
WHERE customer_id = $1
LIMIT 1;

-- name: ListCustomers :many
SELECT *
FROM customer
ORDER BY customer_id
LIMIT $1 OFFSET $2;
