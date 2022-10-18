package users

import (
	"api-gateaway/cmd/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func VerifyToken(token string) models.User {
	jsonStr := []byte(fmt.Sprintf(`{"accessToken" :"%v"}`, token))
	resp, err := http.Post("http://api_user:3000/api/v1/user/token", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err.Error())
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var user models.User
	json.Unmarshal(body, &user)
	return user
}
