package cmd_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	db_test "user/__test__/db"
	"user/controllers"
	"user/models"
)

var baseUrl = "http://localhost:3000/api/v1/user"

func testingApi(method string, path string, payload []byte) string {
	r := httptest.NewRequest(method, baseUrl+path, bytes.NewBuffer(payload))
	r.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	r.Header.Set("Content-Type", "application/json")
	if path == "/login" {
		controllers.Login(w, r)
	} else if path == "/register" {
		controllers.Register(w, r)
	} else if path == "/token" {
		controllers.CheckToken(w, r)
	}
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	return string(data)
}

func TestLogin(t *testing.T) {
	db_test.Connect()
	db_test.Migrate()
	defer db_test.DropTable()

	var payload = []byte(`{
		"username": "Budi",
		"password": "pass"
	}`)

	addUser := testingApi(http.MethodPost, "/register", payload)
	if addUser != `{"message":"Success Register"}` {
		t.Error("Failed Add User")
	}
	var token models.Token

	loginWithoutPayload := testingApi(http.MethodPost, "/login", payload)
	json.Unmarshal([]byte(loginWithoutPayload), &token)
	if token.AccessToken == "" {
		t.Error("Failed login")
	}

	var accessToken = []byte(fmt.Sprintf(`{"accessToken":"%v"}`, token.AccessToken))
	checkToken := testingApi(http.MethodPost, "/token", accessToken)
	if checkToken != `{"userId":1,"username":"Budi"}` {
		t.Error("Invalid token")
	}
}
