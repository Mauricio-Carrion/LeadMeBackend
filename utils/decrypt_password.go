package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func DecryptPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil 
}