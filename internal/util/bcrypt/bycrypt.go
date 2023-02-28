package bcryptUtil

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncryptString(str string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return string(hashedByte), nil
}
