-- +goose Up
-- +goose StatementBegin
CREATE TABLE books (
  id         BIGSERIAL PRIMARY KEY,
  name       TEXT NOT NULL,
  author     TEXT NOT NULL,
  stock      INT DEFAULT 0 NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books;
-- +goose StatementEnd
