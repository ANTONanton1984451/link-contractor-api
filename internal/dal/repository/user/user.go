package user

import (
	"context"
	"link-contractor-api/internal/dal"
	"link-contractor-api/internal/entrypoint"
	"time"
)

type (
	repository struct {
		pool dal.Pool
	}
)

func (repo *repository) InBan(ctx context.Context, user entrypoint.User) (bool, *time.Time, error) {
	return false, nil, nil
}

func (repo *repository) GetByExternalID(ctx context.Context, externalID int64) (entrypoint.User, error) {
	return entrypoint.User{
		ID: 1,
	}, nil
}

func New(pool dal.Pool) *repository {
	return &repository{pool: pool}
}
