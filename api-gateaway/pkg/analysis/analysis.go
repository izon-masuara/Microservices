package analysis

import (
	"api-gateaway/cmd/db"
	"api-gateaway/cmd/models"
	"api-gateaway/pkg/users"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Analysis(r *http.Request) models.Files {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	r.Body.Close()
	var postAnalysis models.PostAnalysis
	json.Unmarshal(body, &postAnalysis)
	token := users.VerifyToken(postAnalysis.AccessToken)
	postAnalysis.UserId = token.UserId
	jsonRequest := fmt.Sprintf(`{
			"user_id": "%v",
			"category": "%v",
			"tags" : "%v",
			"date" : "%v",
			"duration" : "%v",
			"total_duration" : "%v"
		  }`,
		postAnalysis.UserId,
		postAnalysis.Category,
		postAnalysis.Tags,
		postAnalysis.Date,
		postAnalysis.Duration,
		postAnalysis.TotalDuration)
	_, err = http.Post("http://localhost:5000/api/v1/analysis/", "apilcation/json", bytes.NewBuffer([]byte(jsonRequest)))
	if err != nil {
		log.Fatal(err.Error())
	}
	response, err := http.Get("http://localhost:5000/api/v1/analysis/" + fmt.Sprintf("%v", postAnalysis.UserId))
	if err != nil {
		log.Fatal(err.Error())
	}
	recomendation, _ := ioutil.ReadAll(response.Body)
	var msg models.ResponseMessage
	json.Unmarshal(recomendation, &msg)
	client := db.Client
	if msg.Status == 400 {
		dataMovies := client.Get("movies")
		dataFiles := models.Files{}
		json.Unmarshal([]byte(dataMovies.Val()), &dataFiles)
		return dataFiles
	} else if msg.Status == 200 {
		dataMovies := client.Get("movies")
		dataFiles := models.Files{}
		json.Unmarshal([]byte(dataMovies.Val()), &dataFiles)
		newData := models.Files{}
		count := 0
		for _, v := range dataFiles {
			if v.Category == msg.Data.Category && count <= 10 {
				newData = append(newData, v)
				count++
			} else if v.Category != msg.Data.Category && count > 10 {
				newData = append(newData, v)
			}
		}
		res, _ := json.Marshal(newData)
		client.Set(fmt.Sprintf("userId:%v", postAnalysis.UserId), string(res), -1)
		return newData
	} else {
		return models.Files{}
	}
}
