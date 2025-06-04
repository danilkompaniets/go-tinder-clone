-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "postgis";

CREATE TABLE IF NOT EXISTS users_preferences(
    id               uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id          uuid NOT NULL UNIQUE,
    preferred_gender TEXT NOT NULL,
    age_min          INT,
    age_max          INT,
    position         GEOGRAPHY(Point, 4326),
    radius           INT,
    foreign key (user_id) REFERENCES users (id)
);

-- +goose Down
DROP TABLE IF EXISTS users_preferences;