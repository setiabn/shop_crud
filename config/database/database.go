package database

import (
	"shop/config"
	"shop/config/database/migration"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dbAddr := config.Get("DB_ADDR")
	dbUser := config.Get("DB_USER")
	dbPass := config.Get("DB_PASS")
	dbName := config.Get("DB_NAME")

	dbCharset := "utf8mb4"
	dbLoc := "Local"

	// Connect to the MySQL database server
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbAddr + ")/" + "?charset=" + dbCharset + "&parseTime=True&loc=" + dbLoc
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Create a new database if it does not exist
	result := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if result.Error != nil {
		panic("failed to create database: " + result.Error.Error())
	}

	// Connect to the newly created database
	dsn = dbUser + ":" + dbPass + "@tcp(" + dbAddr + ")/" + dbName + "?charset=" + dbCharset + "&parseTime=True&loc=" + dbLoc
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	migration.RunMigration(db)

	return db
}
