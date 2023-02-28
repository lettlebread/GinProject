package jwtUtil

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"crypto/rsa"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

var (
	prvKey *rsa.PrivateKey
	pubKey *rsa.PublicKey
)

type authClaims struct {
	jwt.StandardClaims
	Account string `json:"account`
}

func Init() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	prvByte, err := ioutil.ReadFile(os.Getenv("JWT_PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	prk, err := jwt.ParseRSAPrivateKeyFromPEM(prvByte)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	pubByte, err := ioutil.ReadFile(os.Getenv("JWT_PUBLIC_KEY"))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	puk, err := jwt.ParseRSAPublicKeyFromPEM(pubByte)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	prvKey = prk
	pubKey = puk
}

func CreateToken(account string) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims = &authClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		Account: account,
	}

	tokenString, err := t.SignedString(prvKey)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return tokenString, nil
}

func VarifyToken(str string) (string, error) {
	token, err := jwt.ParseWithClaims(str, &authClaims{}, func(token *jwt.Token) (interface{}, error) {
		return pubKey, nil
	})

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	claims := token.Claims.(*authClaims)

	return claims.Account, nil
}
