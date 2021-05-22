package parsing

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"parse_photo_links/cfg"
	"time"
)

// getHTML - get content from site
func getHTML(url string, cfg *cfg.Config) (string, error) {
	log.Printf("Start <getHTML>, url: %s", url)
	defer log.Printf("End <getHTML>")

	client := &http.Client{Timeout: cfg.App.Timeout * time.Second}

	resp, err := client.Get(url)

	if err != nil {
		return "", fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	log.Printf("GetXML, Status code: %v, Url: %s", resp.StatusCode, url)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Read body: %v", err)
	}

	return string(data), nil
}
