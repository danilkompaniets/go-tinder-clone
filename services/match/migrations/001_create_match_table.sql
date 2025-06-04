-- +goose Up
CREATE TABLE IF NOT EXISTS match(
    from_id       UUID NOT NULL,
    to_id         UUID NOT NULL,
    from_decision BOOLEAN,
    to_decision   BOOLEAN,
    PRIMARY KEY (from_id, to_id),
    CHECK (from_id <> to_id)
);

-- +goose Down
DROP TABLE IF EXISTS match;