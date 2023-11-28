package db

import (
	"fmt"

	"github.com/samvimes01/go-rest1/models"
	"github.com/samvimes01/go-rest1/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {

	var (
		databaseUser     string = utils.GetValue("POSTGRES_USER")
		databasePassword string = utils.GetValue("POSTGRES_PASSWORD")
		databaseHost     string = utils.GetValue("POSTGRES_HOST")
		databasePort     string = utils.GetValue("POSTGRES_PORT")
		databaseName     string = utils.GetValue("POSTGRES_DB")
	)

	var dataSource string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Berlin", databaseHost, databaseUser, databasePassword, databaseName, databasePort)

	var err error

	DB, err = gorm.Open(postgres.Open(dataSource), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to the database")

  DB.AutoMigrate(&models.User{}, &models.Item{})
}