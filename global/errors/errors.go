package err

import "errors"

var (
	ErrConnectDatabase = errors.New("FAILED_CONNECT_TO_DATABASE")

	ErrMigrateDatabase = errors.New("FAILED_MIGRATE_DATABASE")
)
