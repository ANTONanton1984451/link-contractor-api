package dependencies

import (
	"link-contractor-api/internal/dal/pool"
	"time"

	"github.com/jackc/pgx/log/zapadapter"
	"go.uber.org/zap"
)

// InitPool  инициализация пула коннекшенов
func InitPool(zlg *zap.Logger, dbDsn string, maxConn, connectRetriesCount int64, retryWait time.Duration) error {
	logger := zapadapter.NewLogger(zlg)
	return pool.InitPool(logger, dbDsn, maxConn, connectRetriesCount, retryWait)
}
