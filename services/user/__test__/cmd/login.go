package cmd_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func loginTest(t *testing.T) {
	loginWithWrongUsername := testingApi(http.MethodPost, "/login", payloadWrongUsername)
	loginWithWrongPassword := testingApi(http.MethodPost, "/login", payloadWrongPassword)
	validUser := testingApi(http.MethodPost, "/login", payload)
	if loginWithWrongUsername != "Invalid username or password\n" {
		t.Error("Bug cause user login with wrong username or password")
	}
	if loginWithWrongPassword != "Invalid usernname or password\n" {
		fmt.Println(loginWithWrongPassword)
		t.Error("Bug cause user login with wrong username or password")
	}
	json.Unmarshal([]byte(validUser), &token)
	if len(token.AccessToken) <= 0 {
		t.Error("Bug cause router do not return access token")
	}
}
