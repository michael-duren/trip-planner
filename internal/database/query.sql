-- name: GetUser :one
SELECT * FROM "User"
WHERE user_id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM "User"
WHERE email = $1 LIMIT 1;

-- name: ListUserTrips :many
SELECT * FROM TripList
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: CreateUser :one
INSERT INTO "User" 
    (email)
 VALUES (
  $1
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE "User"
  set email = $2
WHERE user_id = $1;

-- name: DeleteUser :exec
DELETE FROM "User"
WHERE user_id = $1;
