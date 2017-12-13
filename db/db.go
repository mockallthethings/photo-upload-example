package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:@/mysql")
}
