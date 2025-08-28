package utils

import (
	"errors"
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
func ValidateToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if ok != true {
			return false, errors.New("unexpected signing method for token creation")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return errors.New("could not parse token")
	}
	tokenValid := parsedToken.Valid
	if !tokenValid {
		return errors.New("token validation failed")
	}
	//claims, ok := parsedToken.Claims.(jwt.MapClaims)
	//if ok != true {
	//	return errors.New("Invalid token claim")
	//}
	//email := claims["email"].(string)
	//id := claims["id"].(int)
	return nil
}
