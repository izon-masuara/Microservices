package files

import (
	"api-gateaway/cmd/db"
	"api-gateaway/cmd/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetFiles() models.Files {
	client := db.Client
	data := client.Get("movies")
	if len(data.Val()) == 0 {
		resp, err := http.Get("http://localhost:3001/api/v1/files/")
		if err != nil {
			log.Fatal(err.Error())
		}
		readData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		resp.Body.Close()
		var dataFiles = models.Files{}
		json.Unmarshal(readData, &dataFiles)
		result, err := json.Marshal(dataFiles)
		if err != nil {
			log.Fatal(err.Error())
		}
		client.Set("movies", string(result), -1)
	}

	dataMovies := client.Get("movies")
	var dataFiles = models.Files{}
	json.Unmarshal([]byte(dataMovies.Val()), &dataFiles)

	return dataFiles
}
