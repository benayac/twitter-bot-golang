package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/twitter-bot-golang/controllers"
)

func main() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	r.HandleFunc("/webhooks/twitter", controllers.Webhook)
	http.ListenAndServe(":"+port, r)
}
