package cmd_test

import (
	"net/http"
	"testing"
)

func addUserTest(t *testing.T) {
	addUser := testingApi(http.MethodPost, "/register", payload)
	addUserWithoutPayload := testingApi(http.MethodPost, "/register", nil)
	addUserWithoutUsername := testingApi(http.MethodPost, "/register", payloadWithoutUsername)
	addUserWithoutPassword := testingApi(http.MethodPost, "/register", payloadWithoutPass)
	addUserAlreadyExits := testingApi(http.MethodPost, "/register", payload2)
	if addUser != `{"message":"Success Register"}` {
		t.Error("Failed Add User")
	}
	if addUserWithoutPayload != `{"message":"Username or Password must be more than 3 characters"}` {
		t.Error("Bug cause user can be adding without payload")
	}
	if addUserWithoutUsername != `{"message":"Username or Password must be more than 3 characters"}` {
		t.Error("Bug cause user can be adding without username")
	}
	if addUserWithoutPassword != `{"message":"Username or Password must be more than 3 characters"}` {
		t.Error("Bug cause user can be adding without password")
	}
	if addUserAlreadyExits != `{"message":"username already exists"}` {
		t.Error("Bug cause database have duplicate user")
	}
}
