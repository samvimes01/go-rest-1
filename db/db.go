package db

import (
	"fmt"

	"github.com/samvimes01/go-rest1/models"
	"github.com/samvimes01/go-rest1/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(dbHost, dbName string) {

	var (
		databaseUser     string = utils.GetValue("DB_USER")
		databasePassword string = utils.GetValue("DB_PASSWORD")
		databasePort     string = utils.GetValue("DB_PORT")
		databaseHost     string = dbHost
		databaseName     string = dbName
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
