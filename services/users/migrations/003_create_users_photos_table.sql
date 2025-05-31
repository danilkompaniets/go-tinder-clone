-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users_photos(
    id       uuid DEFAULT uuid_generate_v4(),
    user_id  uuid NOT NULL,
    url      TEXT NOT NULL,
    position INT  DEFAULT 1
);

-- +goose Down
DROP TABLE IF EXISTS users_photos;