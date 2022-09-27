package helpers

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func GetToken(userId int, username string) (string, error) {
	var signinKey = os.Getenv("JWT_KEY")
	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userId,
		"username": username,
	})
	accessToken, err := tokenString.SignedString([]byte(signinKey))
	return accessToken, err
}

func VerifyToken(tokenString string) (jwt.Claims, error) {
	var signinKey = os.Getenv("JWT_KEY")
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(signinKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
