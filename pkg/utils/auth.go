package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func CompareHashPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
