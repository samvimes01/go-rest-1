package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/samvimes01/go-rest1/db"
	"github.com/samvimes01/go-rest1/models"
	"github.com/samvimes01/go-rest1/utils"
	"golang.org/x/crypto/bcrypt"
)

func Signup(userInput models.UserRequest) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	var user models.User = models.User{
		ID:       uuid.New().String(),
		Email:    userInput.Email,
		Password: string(password),
	}

	db.DB.Create(&user)

	token, err := utils.GenerateNewAccessToken()

	if err != nil {
		return "", err
	}

	return token, nil
}

func Login(userInput models.UserRequest) (string, error) {
	var user models.User

	result := db.DB.Take(&user, "email = ?", userInput.Email) // First is ORDERED BY id which is nonsense for uuid

	if result.RowsAffected == 0 {
		return "", errors.New("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))

	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateNewAccessToken()

	if err != nil {
		return "", err
	}

	return token, nil
}
