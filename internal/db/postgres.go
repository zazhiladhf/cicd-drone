package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB(host, port, user, pass, dbname string) (db *sql.DB, err error) {
	dsn := fmt.Sprintf("host=%v port=%s user=%s password=%v dbname=%v sslmode=disable", host, port, user, pass, dbname)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	return
}
