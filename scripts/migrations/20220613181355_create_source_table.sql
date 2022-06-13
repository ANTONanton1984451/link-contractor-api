-- +goose Up
-- +goose StatementBegin
CREATE TABLE source
(
    id         BIGSERIAL NOT NULL PRIMARY KEY,
    name       TEXT UNIQUE,
    created_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO source (name)
VALUES ('dev'),
       ('vk');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE source;
-- +goose StatementEnd
