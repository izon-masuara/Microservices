package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HassPass(pass string) (interface{}, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(pass), 5)
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
