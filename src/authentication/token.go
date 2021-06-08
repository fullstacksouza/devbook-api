package authentication

import (
	"devbook-api/src/config"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId string) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(config.SecretKey) //secret
}

func ValidToken(r *http.Request) error {
	return nil
}
