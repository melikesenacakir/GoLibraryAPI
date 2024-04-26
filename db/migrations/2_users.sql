-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id       BIGSERIAL PRIMARY KEY,
  role     TEXT CHECK (role IN ('admin','user')) NOT NULL,
  username TEXT NOT NULL,
  password TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
