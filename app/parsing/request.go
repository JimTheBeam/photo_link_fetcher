package parsing

import (
	"log"
	"net/http"
	"parse_photo_links/cfg"
	"time"

	"golang.org/x/net/html"
)

// getHTML - get content from site
func getHTML(url string, cfg *cfg.Config) (*html.Node, error) {
	log.Printf("Start <getHTML>, url: %s", url)
	defer log.Printf("End <getHTML>")

	client := &http.Client{Timeout: cfg.App.Timeout * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		log.Printf("getHTML: %v", err)
		return nil, errGetHTML{link: url}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Status code: %v", resp.StatusCode)
		return nil, errStatusCode{link: url, code: resp.StatusCode}
	}

	// log.Printf("GetXML, Status code: %v, Url: %s", resp.StatusCode, url)

	data, err := html.Parse(resp.Body)
	if err != nil {
		log.Printf("Read body: %v", err)
		return nil, errParseBody{link: url}
	}

	return data, nil
}
