package middleware

import (
	"fmt"
	"time"

	"github.com/dragranzer/Golang-Mini-Project/constant"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{
		"userid": userId,
		"role":   "admin",
		"exp":    time.Now().Add(time.Hour * 1).Unix,
	}
	// claims["userId"] = userId
	// claims["exp"] = time.Now().Add(time.Hour * 1).Unix
	fmt.Println(claims)
	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println("token ", token)
	// fmt.Println("token kali ", (token.SignedString([]byte(constant.SECRET_JWT))))
	token, err := tokenWithClaims.SignedString([]byte(constant.SECRET_JWT))
	fmt.Println("token ", token)
	return token, err
}
