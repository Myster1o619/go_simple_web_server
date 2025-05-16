package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(passwrd string) (string, error) {
	bytePass, err := bcrypt.GenerateFromPassword([]byte(passwrd), bcrypt.DefaultCost)
	return string(bytePass), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
