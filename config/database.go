package config

import (
	"log"

	"github.com/nelsonomoi/usingJWT/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(conncetionString string) {
	Instance, dbError = gorm.Open(mysql.Open(conncetionString),&gorm.Config{})

	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate()  {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}