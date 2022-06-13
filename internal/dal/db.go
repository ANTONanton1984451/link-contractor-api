package dal

import (
	"context"

	"github.com/jackc/pgx"
)

type Pool interface {
	GetConn(ctx context.Context) (*pgx.Conn, error)
}
