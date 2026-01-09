package database

import (
	"os"
	"path/filepath"

	"github.com/BlackMage1T/genshin-pull-prio/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// 1. Use a simple, consistent path relative to your project root
	dbPath := "database/character.db"

	// 2. Ensure the folder exists
	dir := filepath.Dir(dbPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// This creates the "data" folder if it's missing
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic("failed to create database directory: " + err.Error())
		}
	}

	// 3. Open the database using the SAME variable (dbPath)
	// Don't hardcode "database/genshin.db" here
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Auto-Migration
	db.AutoMigrate(
		&models.Character{},
		&models.Talent{},
		&models.UpgradeLevel{},
	)

	return db
}
