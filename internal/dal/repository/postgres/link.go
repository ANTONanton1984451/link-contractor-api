package postgres

import "database/sql"

type (
	repository struct {
		connection sql.Conn
	}
)
