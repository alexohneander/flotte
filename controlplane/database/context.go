package database

import (
	"flotte/controlplane/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var databasePath string = "/tmp/flotte.db"

func InitializeDatabase() *gorm.DB {
	log.Println("Database Context: Initializing Database..")
	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{
		// Silent Database Logging, because it is only used for deep debugging
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("Database Context: failed to connect database")
	}

	// Migrate the schema
	log.Println("Database Context: Migrating Database..")
	db.AutoMigrate(&model.ControlPlane{}, &model.WorkerPlane{})
	log.Println("Database Context: Database initialized")

	return db
}
