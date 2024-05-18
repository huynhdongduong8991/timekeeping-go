package lib

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DBConn *sql.DB

func ConnectDB() error {
	config, err := NewConfig()
	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DB.DB_HOST,
		config.DB.DB_PORT,
		config.DB.DB_USER,
		config.DB.DB_PASSWORD,
		config.DB.DB_NAME,
		config.DB.SSL_MODE,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open a DB connection: %v", err)
	}

	DBConn = db

	return nil
}
