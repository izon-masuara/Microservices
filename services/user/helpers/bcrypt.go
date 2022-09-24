package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HassPass(pass string) string {
	res, err := bcrypt.GenerateFromPassword([]byte(pass), 5)
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func ComparePass(passFromDb string, pass string) (interface{}, error) {
	err := bcrypt.CompareHashAndPassword([]byte(passFromDb), []byte(pass))
	if err != nil {
		return nil, err
	}
	return true, nil
}
