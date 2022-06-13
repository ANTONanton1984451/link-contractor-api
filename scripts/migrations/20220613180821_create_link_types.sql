-- +goose Up
-- +goose StatementBegin
CREATE TYPE create_type AS ENUM ('random','user_generated');
ALTER TABLE link
    ADD COLUMN create_type create_type;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE link
    DROP COLUMN create_type;
DROP TYPE create_type;
-- +goose StatementEnd
