package middleware

import (
	"time"

	"github.com/dragranzer/Golang-Mini-Project/constant"

	jwt "github.com/golang-jwt/jwt"
)

func CreateToken(userId int, name string) (string, error) {
	claims := jwt.MapClaims{
		"userid": userId,
		"name":   name,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	}
	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenWithClaims.SignedString([]byte(constant.SECRET_JWT))
	return token, err
}
