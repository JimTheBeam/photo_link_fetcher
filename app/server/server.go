package server

import (
	"encoding/json"
	"io/ioutil"
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

	// convert r.body to []byte
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Convert r.body to []byte: %v", err)
	}

	pageUrls, err := parsing.ParseAll(&cfg.Config{}, body)
	if err != nil {
		log.Printf("ParseAll: %v", err)
		json.NewEncoder(w).Encode(ResponseJson{
			ErrorCode:    422,
			ErrorMessage: err.Error(),
			Result:       pageUrls,
		},
		)
	} else {
		json.NewEncoder(w).Encode(ResponseJson{Result: pageUrls})
	}
}
