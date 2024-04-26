-- name: BookBorrow :many
INSERT INTO bookBorrows (borrowDate,returnDate,user_id,book_id,status) 
VALUES ($1, $2, $3,$4,$5) RETURNING *;

-- name: GetBorrows :many
SELECT
    bb.id AS borrow_id,
    bb.borrowDate,
    bb.returnDate,
    bb.status,
    u.username AS user_username,
    b.name AS book_name
FROM
    bookBorrows bb
INNER JOIN
    users u ON bb.user_id = u.id
INNER JOIN
    books b ON bb.book_id = b.id
WHERE status= $1 OR bb.returnDate < CURRENT_TIMESTAMP;

-- name: GetBorrowHistory :many
SELECT
    bb.id AS borrow_id,
    bb.borrowDate,
    bb.returnDate,
    bb.status,
    u.username AS user_username,
    b.name AS book_name
FROM
    bookBorrows bb
INNER JOIN
    users u ON bb.user_id = u.id
INNER JOIN
    books b ON bb.book_id = b.id;

-- name: Returnbook :exec 
UPDATE bookBorrows
  set status = $2
WHERE id = $1;