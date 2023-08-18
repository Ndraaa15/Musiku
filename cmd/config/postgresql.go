package config

import (
	"fmt"
	"os"
)

func PostgresqlConfig() string {
	DataSourceName := fmt.Sprintf("%s",
		os.Getenv("A"))

	return DataSourceName
}
