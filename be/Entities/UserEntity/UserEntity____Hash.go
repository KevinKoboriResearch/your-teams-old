package UserEntity

import (
	"golang.org/x/crypto/bcrypt"
)

//-------------------------------- BG --------------------------------//
func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 1)
	return string(bytes), err
}

//-------------------------------- BG --------------------------------//
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
