package database

import (
	"database/sql"
	"log"
	"time"

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

// 监控数据库连接的健康状况，定期Ping数据库并处理异常情况
func monitorDBConnection() {
	ticker := time.NewTicker(5 * time.Minute) // 每隔5分钟检查一次连接健康状况，可调整
	defer ticker.Stop()
	for range ticker.C {
		if DB == nil {
			log.Println("数据库连接对象为nil，可能尚未初始化或已意外关闭，等待重新初始化...")
			continue
		} else {
			log.Println("数据库连接正常...")
		}

		// 启动连接健康检查协程
		go monitorDBConnection()

		err := DB.Ping()
		if err != nil {
			// 连接出现问题，尝试重新连接
			log.Printf("Database connection lost. Attempting to reconnect...")
			reconnect()
		}
	}
}

// 尝试重新连接数据库
func reconnect() {
	var err error
	dbConfig := configs.AppConfig.Database
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Error reopening database: %v", err)
		return
	}

	// 再次检查连接是否可用
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pinging database after reconnect: %v", err)
		return
	}
	log.Println("Database reconnected successfully")
}

// 关闭数据库连接
func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatalf("Error closing database: %v", err)
	}
}
