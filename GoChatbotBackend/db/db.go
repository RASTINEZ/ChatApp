package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	connStr := os.Getenv("DB_CONN_STRING")


	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("DB open error: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("DB ping error: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL")
	return nil
}
