package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connection() (*gorm.DB, error) {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")

	var gormConfig *gorm.Config = &gorm.Config{}

	environment := os.Getenv("APP_ENV")

	if environment != "production" {

		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer (log output)
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level (Info will log all queries)
				IgnoreRecordNotFoundError: true,        // Ignore "record not found" errors
				Colorful:                  true,        // Enable color
			},
		)
		gormConfig = &gorm.Config{
			Logger: newLogger,
		}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)

	if err != nil {
		return nil, err
	}

	return db, nil
}
