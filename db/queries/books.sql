-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: GetBooks :many
SELECT * FROM books;

-- name: UpdateBook :exec
UPDATE books
  set name = $2,
  author = $3,
  stock = $4
WHERE id = $1;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;

-- name: CreateBook :one
INSERT INTO books (name, author,stock) 
VALUES ($1, $2, $3) RETURNING *;