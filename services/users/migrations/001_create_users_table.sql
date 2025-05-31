-- +goose Up
CREATE TABLE IF NOT EXISTS users(
    id         uuid PRIMARY KEY,
    username   TEXT NOT NULL,
    email      TEXT NOT NULL,
    first_name TEXT,
    bio        TEXT,
    gender     TEXT,
    birth_date DATE,
    city       TEXT NOT NULL,
    avatar_url TEXT NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS users;
