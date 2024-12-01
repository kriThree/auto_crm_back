-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);
ALTER TABLE users ADD role_id INTEGER REFERENCES roles(id);
INSERT INTO roles (name) VALUES ('admin');
INSERT INTO roles (name) VALUES ('client');
INSERT INTO roles (name) VALUES ('owner');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM roles WHERE name IN ('admin', 'client', 'owner');
ALTER TABLE users DROP COLUMN role_id;
-- +goose StatementEnd
