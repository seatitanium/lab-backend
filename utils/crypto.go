package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func VerifyHash(hash []byte, possiblePassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, possiblePassword)
	return err == nil
}
