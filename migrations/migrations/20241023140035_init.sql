-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULl,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
create TABLE owners (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
create TABLE admins (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE TABLE autoservices (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    owner_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (owner_id) REFERENCES owners(id)
);
CREATE TABLE cars (
    id SERIAL PRIMARY KEY,
    number VARCHAR(255) NOT NULL UNIQUE,
    description VARCHAR(255),
    client_id INTEGER NOT NULL,
    FOREIGN KEY (client_id) REFERENCES clients(id)
);
create TABLE catalogs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    creater_id INTEGER NOT NULL,
    FOREIGN KEY (creater_id) REFERENCES admins(id)
);
create TABLE works (
    id SERIAL PRIMARY KEY,
    cost INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    catalog_id INTEGER NOT NULL,
    FOREIGN KEY (catalog_id) REFERENCES catalogs(id)
);
create TABLE operations (
    id SERIAL PRIMARY KEY,
    description VARCHAR(255) NOT NULL,

    car_id INTEGER NOT NULL,
    work_id INTEGER NOT NULL,
    autoservice_id INTEGER NOT NULL,

    FOREIGN KEY (car_id) REFERENCES cars(id),
    FOREIGN KEY (work_id) REFERENCES works(id),
    FOREIGN KEY (autoservice_id) REFERENCES autoservices(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE operations;
DROP TABLE works;
DROP TABLE catalogs;
DROP TABLE cars;
DROP TABLE autoservices;
DROP TABLE clients;
DROP TABLE owners;
DROP TABLE admins;
DROP TABLE users;
-- +goose StatementEnd
