package jsonparse

import (
	"encoding/json"
	"log"

	"github.com/go-playground/validator"
)

// ParseJSON - parse an incoming json
func ParseJSON(body []byte, urls *IncomingJSON) error {

	// parse urls from incoming json and put them in urls
	err := json.Unmarshal(body, &urls)
	if err != nil {
		log.Printf("Unmarshal JSON. %v\n", err)
		return ErrParseJson
	}
	log.Printf("<ParseJSON> result: %v\n", urls)

	// parse urls from incoming json and put them in urls
	// if err := json.NewDecoder(body).Decode(&urls); err != nil {
	// 	log.Printf("Decoding json: %v", err)
	// 	return ErrParseJson
	// }

	// validate json
	v := validator.New()

	if err := v.Struct(urls); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			log.Printf("validate json: %v\n", e)
		}
		return ErrParseJson
	}
	return nil
}
