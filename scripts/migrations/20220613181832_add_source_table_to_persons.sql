-- +goose Up
-- +goose StatementBegin
ALTER TABLE person
    ADD COLUMN source_id BIGINT REFERENCES source;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE person
    DROP COLUMN source_id;
-- +goose StatementEnd
