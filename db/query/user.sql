-- name: CreateUser :exec
INSERT INTO users (username,password,role,active)
VALUES ($1,$2,$3,$4);

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: UpdateUser :exec
UPDATE users
SET password = $2,
    role = $3,
    active = $4
WHERE username = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;