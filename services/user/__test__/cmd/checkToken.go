package cmd_test

import (
	"fmt"
	"net/http"
	"testing"
)

func accessTokenTest(t *testing.T) {
	var accessToken = []byte(fmt.Sprintf(`{"accessToken":"%v"}`, token.AccessToken))
	var invalidAccessToken = []byte(`{"accessToken":"asdlkshadieran97845kljhlaw"}`)
	checkToken := testingApi(http.MethodPost, "/token", accessToken)
	checkInvalidToken := testingApi(http.MethodPost, "/token", invalidAccessToken)
	if checkToken != `{"userId":1,"username":"Budi"}` {
		t.Error("Invalid token")
	}
	if checkInvalidToken != `{"message":"Bad request"}` {
		t.Error("Bug cause invalid token return data")
	}
}
