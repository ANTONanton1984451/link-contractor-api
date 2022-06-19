package user

import (
	"link-contractor-api/internal/dal"
)

type (
	repository struct {
		pool dal.Pool
	}
)

func New(pool dal.Pool) *repository {
	return &repository{pool: pool}
}
