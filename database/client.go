package database

import (
	config "dataViewer/config"
	entities "dataViewer/entities"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var errsql error

func OpenConnection() {
	db, err := gorm.Open(postgres.Open(config.ConnectionString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to Database %s", db.Name())
}

func Migrate() {
	Instance.AutoMigrate(&entities.User{})
	log.Println("database migration completed")
}
