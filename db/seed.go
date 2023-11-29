package db

import (
	"errors"
	"fmt"

	"github.com/samvimes01/go-rest1/models"
	"github.com/samvimes01/go-rest1/utils"
	"golang.org/x/crypto/bcrypt"
)

// CleanSeeders performs clean up mechanism after testing
func CleanSeeders() {
	itemResult := DB.Exec("TRUNCATE items")
	userResult := DB.Exec("TRUNCATE users")

	var isFailed bool = itemResult.Error != nil || userResult.Error != nil

	if isFailed {
		panic(errors.New("error when cleaning up seeders"))
	}

	fmt.Println("Seeders are cleaned up successfully")
}

// SeedItem returns recently created items from the database
func SeedItem() (models.Item, error) {
	item, err := utils.CreateFaker[models.Item]()
	if err != nil {
		return models.Item{}, nil
	}

	DB.Create(&item)
	fmt.Println("Item seeded to the database")

	return item, nil
}

// SeedUser returns recently created user from the database
func SeedUser() (models.User, error) {
	user, err := utils.CreateFaker[models.User]()
	if err != nil {
		return models.User{}, err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, err
	}

	var inputUser models.User = models.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: string(password),
	}

	DB.Create(&inputUser)
	fmt.Println("User seeded to the database")

	return user, nil
}
