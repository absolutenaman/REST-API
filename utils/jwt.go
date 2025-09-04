package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "6ccb75192751c56d05f870c92b4b644548286fd0edd72d41e98c5b1d7710008e"

func TokenGeneration(email string, id int64) (string, error) {
	fmt.Println("!!! email string, id int64", email, id)
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
func ValidateToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method for token creation")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, errors.New("could not parse token")
	}
	if !parsedToken.Valid {
		return 0, errors.New("token validation failed")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	fmt.Println("!!! All claims:", claims)

	idFloat, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("id not found or invalid type in claims")
	}
	fmt.Println("!!! idFloat", int64(idFloat))
	return int64(idFloat), nil
}
