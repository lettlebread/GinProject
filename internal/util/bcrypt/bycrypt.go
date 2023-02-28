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

func ComparePassword(hashedPwd string, pwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	if err != nil {
		return err
	}

	return nil
}
