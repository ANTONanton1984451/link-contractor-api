package pool

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx"
)

var _connPool *pgx.ConnPool

type pgxPool struct{}

func InitPool(logger pgx.Logger, dbDsn string, maxConn, connectRetriesCount int64, retryWait time.Duration) error {
	connCfg, err := pgx.ParseDSN(dbDsn)
	if err != nil {
		return fmt.Errorf("parse dsn: %w", err)
	}

	connCfg.Logger = logger
	poolCfg := pgx.ConnPoolConfig{
		MaxConnections: int(maxConn),
		ConnConfig:     connCfg,
	}

	pool, err := tryConnect(poolCfg, retryWait, int(connectRetriesCount))
	if err != nil {
		return fmt.Errorf("create new conn pool: %w", err)
	}

	_connPool = pool
	return nil
}

func tryConnect(config pgx.ConnPoolConfig, tryEvery time.Duration, retriesCount int) (*pgx.ConnPool, error) {
	var (
		err  error
		pool *pgx.ConnPool
	)

	for i := 0; i <= retriesCount; i++ {
		pool, err = pgx.NewConnPool(config)
		if err == nil {
			return pool, nil
		}
		time.Sleep(tryEvery)
	}

	return nil, err
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
