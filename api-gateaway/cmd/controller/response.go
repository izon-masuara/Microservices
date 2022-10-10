package controller

import (
	"api-gateaway/cmd/helper"
	"api-gateaway/cmd/models"
	"api-gateaway/pkg/analysis"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var GetDataFormQuery = func(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	mutation := r.URL.Query().Get("mutation")
	if len(query) != 0 {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		r.Body.Close()
		var accessToken models.Token
		json.Unmarshal(body, &accessToken)
		if len(query) == 0 {
			fmt.Println("error")
		}
		data := helper.ExecuteQuery(query, accessToken.AccessToken)
		res, _ := json.Marshal(data)
		w.Write(res)
	} else if len(mutation) != 0 {
		result := analysis.Analysis(r)
		resp, _ := json.Marshal(result)
		w.Write(resp)
	}

}
