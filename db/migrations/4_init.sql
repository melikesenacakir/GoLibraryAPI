-- +goose Up
INSERT INTO users (role,username,password) VALUES ('admin','admin121','$2a$12$FlqT84U/crYjC0WSj6w9juirybfxfbJ0TpZP9vX0lnWXccMyC8IMG');

-- +goose Down

