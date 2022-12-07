package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/**
 * @Author ANYARONKE Dare Samuel
 */

// DbInstance : database pointer
type DbInstance struct {
	Db *gorm.DB
}

// Database : object instance
var Database DbInstance

// ConnectDb : Connection to the database
func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("dbbneeft.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	log.Println("Connected Successfully to Database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	// Launch of the database migration
	RunMigration(db)
	Database = DbInstance{Db: db}
}
