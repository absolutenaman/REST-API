package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "6ccb75192751c56d05f870c92b4b644548286fd0edd72d41e98c5b1d7710008e"

func TokenGeneration(email string, id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})
	signedString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedString, nil
}
