package database

import (
	"assignment_2/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "RajaSonang1999*"
	dbPort   = "5432"
	dbName   = "assignment2"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})
}

func GetDB() *gorm.DB {
	return db
}
