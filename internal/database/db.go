package database

import (
	"fmt"
	"log"
	"sync"
	"github.com/kianaw22/birthy/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

// InitDB initializes the database connection once using Singleton pattern
func InitDB() {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.DBConfig.Host, config.DBConfig.Port, config.DBConfig.User, config.DBConfig.Password, config.DBConfig.DBName, config.DBConfig.SSLMode,
		)

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("❗️ Error connecting to database: %v", err)
		}

		_ , err = conn.DB()
		if err != nil {
			log.Fatalf("❗️ Error accessing sql.DB from GORM: %v", err)
		}


		db = conn
		log.Println("✅ Database connected successfully on localhost.")
	})
}

// GetDB returns the initialized *gorm.DB instance
func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("❗️ Database is not initialized. Call InitDB() first.")
	}
	return db
}

// CloseDB properly closes the database connection
func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("⚠️ Error while closing database: %v", err)
		return
	}
	sqlDB.Close()
	log.Println("✅ Database connection closed.")
}
