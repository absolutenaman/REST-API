package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPasswordInByteForm, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}
	return string(hashedPasswordInByteForm), nil
}
func ValidatePassword(hashedPassword, userEnteredPassword []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hashedPassword, userEnteredPassword)
	if err != nil {
		return false, err
	}
	return true, nil
}
