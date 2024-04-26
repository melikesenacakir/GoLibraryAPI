-- +goose Up
-- +goose StatementBegin
CREATE TABLE bookBorrows (
  id         BIGSERIAL PRIMARY KEY,
  borrowDate TIMESTAMP NOT NULL,
  returnDate TIMESTAMP NOT NULL,
  user_id    BIGINT NOT NULL,
  book_id    BIGINT NOT NULL,
  status TEXT CHECK (status IN('returned','borrowed')) NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (book_id) REFERENCES books(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE bookBorrows;
-- +goose StatementEnd
