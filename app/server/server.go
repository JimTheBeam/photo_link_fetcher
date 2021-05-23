package server

import (
	"encoding/json"
	"log"
	"net/http"
	"parse_photo_links/app/parsing"
	jsonparse "parse_photo_links/app/parsing/json_parse"
	"parse_photo_links/cfg"

	"github.com/gorilla/mux"
)

func Server() {
	log.Printf("Server is starting")
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/fetch", getLinks).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// TODO: HERE !!!!!!!
func getLinks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var urls jsonparse.IncomingJSON

	if err := json.NewDecoder(r.Body).Decode(&urls); err != nil {
		log.Printf("Decoding json: %v", err)
		// TODO: return!!!!!!!

	}

	pageUrls, err := parsing.ParseAll(&cfg.Config{}, urls)
	if err != nil {
		log.Printf("ParseAll: %v", err)
		// TODO: return!!!!!!!
	}

	json.NewEncoder(w).Encode(pageUrls)
}
