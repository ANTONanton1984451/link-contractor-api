-- +goose Up
-- +goose StatementBegin
CREATE TABLE person
(
    id          BIGSERIAL NOT NULL PRIMARY KEY,
    name        TEXT,
    surname     TEXT,
    external_id TEXT NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE person;
-- +goose StatementEnd
