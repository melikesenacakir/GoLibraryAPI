-- name: CreateUser :one
INSERT INTO users (role,username,password) 
VALUES ($1, $2, $3) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: UpdateUser :exec
UPDATE users
  set username = $2,
  password=$3
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: Login :one
SELECT role, username, password FROM users WHERE username = $1;




