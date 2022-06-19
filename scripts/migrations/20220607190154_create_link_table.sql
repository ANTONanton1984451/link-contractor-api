-- +goose Up
-- +goose StatementBegin
CREATE TABLE link
(
    id          BIGSERIAL NOT NULL PRIMARY KEY,
    path        TEXT      NOT NULL,
    redirect_to TEXT      NOT NULL,
    active      bool      NOT NULL DEFAULT true,
    owner_id    BIGINT
        CONSTRAINT owner_id_fk REFERENCES person ON DELETE CASCADE,
    created_at  TIMESTAMP          DEFAULT NOW(),
    updated_at  TIMESTAMP          DEFAULT NOW()
);


ALTER TABLE link ADD CONSTRAINT  unique_path UNIQUE (path);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE link;
-- +goose StatementEnd
