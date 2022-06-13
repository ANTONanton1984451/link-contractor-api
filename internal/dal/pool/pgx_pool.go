package pool

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx"
)

var _connPool *pgx.ConnPool

type pgxPool struct{}

func InitPool(logger pgx.Logger, dbDsn string, maxConn int64) error {
	connCfg, err := pgx.ParseDSN(dbDsn)
	if err != nil {
		return fmt.Errorf("parse dsn: %w", err)
	}
	connCfg.Logger = logger
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		MaxConnections: int(maxConn),
		ConnConfig:     connCfg,
	})
	if err != nil {
		return fmt.Errorf("create new conn pool: %w", err)
	}

	_connPool = pool
	return nil
}

func (p pgxPool) GetConn(ctx context.Context) (*pgx.Conn, error) {
	if _connPool == nil {
		return nil, errors.New("pool is not init")
	}
	return _connPool.AcquireEx(ctx)
}

func GetPool() pgxPool {
	return pgxPool{}
}
