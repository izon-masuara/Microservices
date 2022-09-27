package helpers

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HassPass(pass string) (interface{}, error) {
	key := os.Getenv("HASH")
	number, err := strconv.Atoi(key)
	if err != nil {
		return nil, err
	}
	res, err := bcrypt.GenerateFromPassword([]byte(pass), number)
	if err != nil {
		return nil, err
	}
	return string(res), nil
}

func ComparePass(passFromDb string, pass string) (interface{}, error) {
	err := bcrypt.CompareHashAndPassword([]byte(passFromDb), []byte(pass))
	if err != nil {
		return nil, err
	}
	return true, nil
}
