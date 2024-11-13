package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // PostgreSQL 驱动，您可以根据需要更换为其他驱动
	// "gorm.io/gorm"

	"backend/configs"
	"fmt"
)

var DB *sql.DB
var err error

func InitDB() error {
	dbConfig := configs.AppConfig.Database
	// Set up the database source string.
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	// Create a database handle and open a connection pool.
	DB, err = sql.Open("mysql", dataSourceName)
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

	log.Println("数据库连接成功")
	return nil
}

// 关闭数据库连接
func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatalf("Error closing database: %v", err)
	}
}
