package dbHelper

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DB"))

	if err != nil {
		return nil, err
	}

	return db, nil
}
