package database

import (
	"flotte/controlplane/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var databasePath string = "/tmp/flotte.db"

func InitializeDatabase() *gorm.DB {
	log.Println("Database Context: Initializing Database..")
	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	log.Println("Database Context: Migrating Database..")
	db.AutoMigrate(&model.ControlPlane{})
	log.Println("Database Context: Database initialized")

	return db
}
