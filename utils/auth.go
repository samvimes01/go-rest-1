package utils

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateNewAccessToken() (string, error) { // JWT token
	secret := GetValue("JWT_SECRET_KEY")

	minutesCount, _ := strconv.Atoi(GetValue("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	claims := jwt.MapClaims{}

	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
