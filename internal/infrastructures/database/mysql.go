package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var _MaxAttempts = 10

var DB *sql.DB

func NewDBMySQL(dsn string) (*sql.DB, error) {
	var err error

	if DB == nil {
		// Attempting to connect with database
		for i := 1; i <= _MaxAttempts; i++ {
			log.Printf("[INFO] \t: Attempting to connect with database.... [%d]", i)

			DB, _ = sql.Open("mysql", dsn)

			if DB.Ping() == nil {
				log.Printf("[SUCCESS] \t: Successfully connected to database")
				break
			}

			time.Sleep(time.Millisecond * 1500)
		}

		// failed to connect with database
		if DB.Ping() != nil {
			log.Fatalf("[FAILED] \t: Failed to connect with database")
		}

	}

	return DB, err
}
