package link

import (
	"context"
	"link-contractor-api/internal/dal"
)

type (
	Redirect interface {
		Store(ctx context.Context, path string, redirectTo string) error
	}

	repository struct {
		pool dal.Pool

		redirect Redirect
	}
)

func New(pool dal.Pool, rd Redirect) *repository {
	return &repository{pool: pool, redirect: rd}
}
