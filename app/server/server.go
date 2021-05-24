package server

import (
	"encoding/json"
	"log"
	"net/http"
	"parse_photo_links/app/parsing"
	"parse_photo_links/cfg"

	"github.com/gorilla/mux"
)

func Server() {
	log.Printf("Server is starting")
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/fetch", getLinks).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

//
func getLinks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pageUrls, err := parsing.ParseAll(&cfg.Config{}, r.Body)
	if err != nil {
		log.Printf("ParseAll: %v", err)
		json.NewEncoder(w).Encode(ResponseJson{Error: err.Error(), Data: pageUrls})
	} else {
		json.NewEncoder(w).Encode(ResponseJson{Data: pageUrls})
	}
}
