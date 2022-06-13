-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.link
(
    id          BIGSERIAL NOT NULL PRIMARY KEY,
    path        TEXT      NOT NULL,
    redirect_to TEXT      NOT NULL,
    active      bool      NOT NULL DEFAULT true,
    owner_id    BIGINT
        CONSTRAINT owner_id_fk REFERENCES person ON DELETE restrict,
    created_at  TIMESTAMP          DEFAULT NOW(),
    updated_at  TIMESTAMP          DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
