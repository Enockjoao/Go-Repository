package config

import (
	"os"

	"github.com/Enockjoao/Go-Repository/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// Check if the database files exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		// Create database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("sqlite directory creation error: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("sqlite file creation error: %v", err)
			return nil, err
		}
		file.Close()

	}

	// Create DB and Connect
	db, err := gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}
	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}
