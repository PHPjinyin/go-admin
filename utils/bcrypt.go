package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func BcryptDecode(b []byte) string {
	hash, err := bcrypt.GenerateFromPassword(b, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func BcryptCheck(password []byte, hashPwd string) bool {
	byteHash := []byte(hashPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, password)
	if err != nil {
		return false
	}
	return true
}
