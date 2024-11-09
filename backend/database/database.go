package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // PostgreSQL 驱动，您可以根据需要更换为其他驱动

	"backend/config"
	"fmt"
)

var DB *sql.DB

func InitDB() error {
	dbConfig := config.AppConfig.Database
	// Set up the database source string.
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	// Create a database handle and open a connection pool.
	DB, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		return err
	}

	// Check if our connection is alive.
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
		return err
	}

	log.Println("Database connected successfully")
	return nil
}

// 关闭数据库连接
func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatalf("Error closing database: %v", err)
	}
}
