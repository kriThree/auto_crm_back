-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULl,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE TABLE autoservices (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    owner_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (owner_id) REFERENCES users(id)
);
CREATE TABLE cars (
    id SERIAL PRIMARY KEY,
    number VARCHAR(255) NOT NULL UNIQUE,
    description VARCHAR(255) 
);
create TABLE works (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);
create TABLE operations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,

    car_id INTEGER NOT NULL,
    work_id INTEGER NOT NULL,
    autoservice_id INTEGER NOT NULL,

    FOREIGN KEY (car_id) REFERENCES users(id),
    FOREIGN KEY (work_id) REFERENCES works(id),
    FOREIGN KEY (autoservice_id) REFERENCES autoservices(id)

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE operations;
DROP TABLE works;
DROP TABLE cars;
DROP TABLE autoservices;
DROP TABLE roles;
DROP TABLE users;
-- +goose StatementEnd
