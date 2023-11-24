package db

import (
	"sesi-11/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnectionPostgres(t *testing.T) {
	host := config.GetEnv("DB_HOST", "localhost")
	port := config.GetEnv("DB_PORT", "5432")
	user := config.GetEnv("DB_USER", "postgres")
	pass := config.GetEnv("DB_PASS", "mysecretpassword")
	dbname := config.GetEnv("DB_NAME", "postgres")

	t.Run("success connect", func(t *testing.T) {
		db, err := ConnectDB(
			host,
			port,
			user,
			pass,
			dbname,
		)
		require.Nil(t, err)
		require.NotNil(t, db)
	})

	t.Run("invalid password", func(t *testing.T) {
		_, err := ConnectDB(
			host,
			port,
			user,
			"invalid-pass",
			dbname,
		)
		require.NotNil(t, err)
	})
}
