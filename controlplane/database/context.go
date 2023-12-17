package database

import (
	"flotte/controlplane/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var databasePath string = "/tmp/flotte.db"

func InitializeDatabase() *gorm.DB {
	log.Println("DatabaseContext: Initializing Database..")
	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		panic("DatabaseContext: failed to connect database")
	}

	// Migrate the schema
	log.Println("DatabaseContext: Migrating Database..")
	db.AutoMigrate(&model.ControlPlane{})
	log.Println("DatabaseContext: Database initialized")

	return db
}
