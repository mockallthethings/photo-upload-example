package db

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

func OpenDB() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   "root",
		Net:    "tcp",
		Addr:   "mysql:3306",
		DBName: "mysql",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	// docker-compose does not support waiting for the db
	// container to be ready, so we need to implement the wait
	// here
	nRetries := 3
	for retryIdx := 0; retryIdx < nRetries; retryIdx += 1 {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(30 * time.Second)
	}
	return db, err
}
