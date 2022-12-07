package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
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

const DNS string = "bbce14802e2bf7:765f916b@tcp(eu-cdbr-west-03.cleardb.net:3306)/heroku_9a62a0aad140d17?charset=utf8mb4&parseTime=True&loc=Local"

// ConnectDb : Connection to the database
func ConnectDb() {
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})

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
