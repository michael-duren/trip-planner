-- User related queries
-- name: GetUser :one
SELECT * FROM "Users"
WHERE user_id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM "Users"
WHERE email = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM "Users"
WHERE username = $1 LIMIT 1;

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


-- TripQueries

-- name: GetTripByName :one
SELECT * FROM Trips
WHERE name = $1;

-- name: CreateTrip :one
INSERT INTO Trips
(user_id, name, created_at)
VALUES
(
    $1,
    $2,
    NOW()
    )
RETURNING *;

-- name: UpdateTripName :exec
UPDATE Trips
SET name = $1
WHERE trip_id = $2;

-- name: DeleteTrip :exec
DELETE FROM Trips
WHERE trip_id = $1;

-- name: ListUserTrips :many
SELECT * FROM Trips
WHERE user_id = $1
ORDER BY created_at DESC;
