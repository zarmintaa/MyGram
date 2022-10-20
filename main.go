package main

import (
	"final-project/database"
	"final-project/routes"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func main() {
	database.StartDB()
	PORT := ":3000"
	err = routes.StartServer(db).Run(PORT)
	if err != nil {
		panic(err)
	}
}
