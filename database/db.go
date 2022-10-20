package database

import (
	"final-project/models"
	"final-project/routes"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "final-project"
	db       *gorm.DB
	PORT     = ":3000"

	err error
)

func StartDB() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connect database", err)
	}

	err := db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	if err != nil {
		log.Fatal("Error migrate database", err)
	}

	err = routes.StartServer(db).Run(PORT)
	if err != nil {
		panic(err)
	}
}
