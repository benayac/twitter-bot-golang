package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/twitter-bot-golang/api"
	"github.com/twitter-bot-golang/helpers"
	"github.com/twitter-bot-golang/models"
)

//Webhook controller
func Webhook(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data := r.URL.Query().Get("crc_token")
		sha := helpers.CreateCRCToken(data)
		response := models.ResponseToken{Response: "sha256=" + sha}
		responseJSON, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(responseJSON))
	case "POST":
		var activity models.ActivityEvent
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}
		// fmt.Println(string(body))
		json.Unmarshal(body, &activity)
		api.CheckActivity(&activity)
	}
}
