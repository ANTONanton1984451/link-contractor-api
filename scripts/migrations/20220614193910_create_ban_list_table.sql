-- +goose Up
-- +goose StatementBegin

CREATE TABLE ban_list (
    id BIGSERIAL PRIMARY KEY,
    person_id BIGINT UNIQUE REFERENCES person ON DELETE CASCADE,
    ban_active BOOL DEFAULT TRUE,
    until TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ban_list;
-- +goose StatementEnd
