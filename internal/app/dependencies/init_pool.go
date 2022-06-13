package dependencies

import (
	"link-contractor-api/internal/dal/pool"

	"github.com/jackc/pgx/log/zapadapter"
	"go.uber.org/zap"
)

func InitPool(zlg *zap.Logger, dbDsn string, maxConn int64) error {
	logger := zapadapter.NewLogger(zlg)
	return pool.InitPool(logger, dbDsn, maxConn)
}
