-- name: GetUser :one
SELECT * FROM "Users"
WHERE user_id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM "Users"
WHERE email = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM "Users"
WHERE username = $1 LIMIT 1;

-- name: ListUserTrips :many
SELECT * FROM TripLists
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: RegisterUser :one
INSERT INTO "Users" 
    (
    email,
    username,
    password
    )
 VALUES (
  $1,$2,$3
)
RETURNING *;

-- name: LoginUser :one
SELECT * FROM "Users"
WHERE email = $1
AND password = $2;

-- name: UpdateUserEmail :exec
UPDATE "Users"
SET email = $1
WHERE user_id = $2;

-- name: UpdateUserPassword :exec
UPDATE "Users"
SET password = $1
WHERE user_id = $2;

-- name: DeleteUsers :exec
DELETE FROM "Users"
WHERE user_id = $1;
