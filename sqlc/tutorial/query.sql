-- name: GetActor :one
SELECT * FROM actor
WHERE actor_id = $1 LIMIT 1;

-- name: ListAllActors :many
SELECT * FROM actor
ORDER BY actor_id;

-- name: ListActors :many
SELECT * FROM actor
ORDER BY actor_id
LIMIT $1 OFFSET $2;