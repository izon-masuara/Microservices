package helper

import (
	"api-gateaway/cmd/db"
	"api-gateaway/cmd/models"
	"api-gateaway/pkg/files"
	"api-gateaway/pkg/users"
	"encoding/json"
	"fmt"
)

var ExecuteQuery = func(query string, token string) models.Files {
	client := db.Client
	if len(token) == 0 {
		data := files.GetFiles()
		return data
	} else {
		user := users.VerifyToken(token)
		data := client.Get(fmt.Sprintf("userId:%v", user.UserId))
		if len(data.Val()) == 0 || data.Val() == "[]" {
			data := files.GetFiles()
			return data
		} else {
			var dataFilesId = models.Files{}
			json.Unmarshal([]byte(data.Val()), &dataFilesId)
			return dataFilesId
		}
	}
}
