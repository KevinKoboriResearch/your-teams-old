package UserEntity

import (
	"golang.org/x/crypto/bcrypt"
)

//__ GENERATE HASH PASSWORD __________________________________________________//
func GenerateHashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 1)

	return string(bytes), err
}

//__ CHECK HASH PASSWORD _____________________________________________________//
func CheckPasswordHash(password string, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
