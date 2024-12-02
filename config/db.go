package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	connStr := "user=postgres dbname=stock_exchange_go password=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// applyMigrations(db)

	return db
}

// func applyMigrations(db *sql.DB) {
// 	migrationFiles := []string{
// 		"./migrations/01_create_users_table.sql",
// 		"./migrations/02_create_stocks_table.sql",
// 		"./migrations/03_create_transactions_table.sql",
// 	}

// 	for _, file := range migrationFiles {
// 		log.Printf("Checking migration file path: %s", file)

// 		data, err := os.ReadFile(file)
// 		if err != nil {
// 			log.Fatalf("Error reading migration file %s: %v", file, err)
// 		}

// 		_, err = db.Exec(string(data))
// 		if err != nil {
// 			log.Fatalf("Error executing migration %s: %v", file, err)
// 		}
// 	}

// }
