package service

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(uuid string) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 2).Unix()
	permissions["userId"] = uuid
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte("secret"))

}
