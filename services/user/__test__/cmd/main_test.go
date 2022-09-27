package cmd_test

import (
	"bytes"
	"io/ioutil"
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

var payload = []byte(`{
	"username": "Budi",
	"password": "pass"
}`)

var payloadWithoutUsername = []byte(`{
	"password": "pass"
}`)

var payloadWithoutPass = []byte(`{"username":"Budi"}`)

var payload2 = []byte(`{
	"username": "Budi",
	"password": "pass"
}`)

var payloadWrongUsername = []byte(`{
	"username": "Budi1",
	"password": "pass"
}`)

var payloadWrongPassword = []byte(`{
	"username": "Budi",
	"password": "passsad"
}`)

var token models.Token

func TestLogin(t *testing.T) {
	db_test.Connect()
	db_test.Migrate()
	defer db_test.DropTable()

	addUserTest(t)
	loginTest(t)
	accessTokenTest(t)
}
