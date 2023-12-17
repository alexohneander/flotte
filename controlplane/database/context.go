package database

import (
	"log"
	"os"
)

var databasePath string = "/tmp/flotte.db"

func InitializeDatabase() {
	// Check if database exists
	if _, err := os.Stat(databasePath); err == nil {
		log.Println("Database: Database already initialized")
		// If yes, read database
	} else if os.IsNotExist(err) {
		log.Println("Database: Database file does not exist")
		// If not, create database
		createDatabase()
		log.Println("Database: Database was created")
	}
}

func createDatabase() {
	file, err := os.Create(databasePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("Database: Database file was created")

	// sqliteDatabase, _ := sql.Open("sqlite3", databasePath)
	// defer sqliteDatabase.Close()
}
