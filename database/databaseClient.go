package database

import (
	config "dataViewer/config"
	"dataViewer/entities"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	Db, err := gorm.Open(postgres.Open(config.ConnectionString), &gorm.Config{})

	if err != nil {
		log.Panic(err)
		return nil
	}
	log.Println("Connected")

	Db.AutoMigrate(entities.User{})
	fmt.Println("Query excuted")
	return Db
}
