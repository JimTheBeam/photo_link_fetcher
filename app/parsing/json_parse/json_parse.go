package jsonparse

import (
	"encoding/json"
	"log"
)

// ParseJSON - parse incoming json
func ParseJSON(JSON string, urls *IncomingJSON) error {

	err := json.Unmarshal([]byte(JSON), &urls)
	if err != nil {
		log.Printf("Error. Unable to unmarshal JSON. %v", err)
		return err
	}
	log.Printf("<ParseJSON> result: %v\n", urls)

	return nil
}
