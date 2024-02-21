package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func VerifyHash(possiblePassword []byte, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(possiblePassword, hash)
	return err == nil
}
