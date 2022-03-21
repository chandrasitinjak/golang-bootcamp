package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	_struct "sesi_8/paractice/struct"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "RajaSonang1999*"
	dbPort   = "5432"
	dbName   = "paractice-restApi"
	db       *gorm.DB
	err      error
)

func StartDB() {
	//paractice-restApi
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	//db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}

	db.Debug().AutoMigrate(_struct.Person{})

}

func GetDB() *gorm.DB {
	return db
}
