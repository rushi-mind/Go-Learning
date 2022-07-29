package utility

import (
	responseTypes "SMT/types/responses"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GetEncryptedPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 9)
	if err != nil {
		log.Fatal(responseTypes.PASSWORD_HASHING_FAILED)
	}
	return string(hashed)
}

func ValidatePassword(password, encryptedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	return err == nil
}
