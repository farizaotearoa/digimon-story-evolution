package utils

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	if GlobalConfig == nil {
		Logger.Fatal("Configuration is not loaded")
	}

	dbUser := GlobalConfig.Database.Username
	dbPassword := GlobalConfig.Database.Password
	dbHost := GlobalConfig.Database.Host
	dbPort := GlobalConfig.Database.Port
	dbName := GlobalConfig.Database.Name
	dbSslMode := GlobalConfig.Database.SSLMode

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSslMode)

	database, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		Logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	DB = database
}
