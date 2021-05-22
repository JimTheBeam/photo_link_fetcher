package parsing

import (
	"fmt"
	"log"
	"net/url"
	jsonparse "parse_photo_links/app/parsing/json_parse"
	"parse_photo_links/cfg"
	"regexp"
	"strings"
)

// var JSON = `{"url":["abc.com","facebook.com","http://www.google.com/", "https://mail.ru/"]}`

var JSON = `{"url":["https://www.mail.ru"]}`

func Parse(cfg *cfg.Config) error {
	var (
		urls    jsonparse.IncomingJSON
		content string
	)

	// parse incoming json
	err := jsonparse.ParseJSON(JSON, &urls)
	if err != nil {
		log.Printf("Parse json: %v", err)
	}

	// get page data:
	for _, oneUrl := range urls.Url {
		// Parse URL
		urlToGet, err := url.Parse(correctUrl(oneUrl))
		if err != nil {
			log.Printf("Parse url: %v\n", err)
			// TODO: Return!!!!!
			return err
		}

		// get page content
		content = parsePage(urlToGet, cfg)

		if err != nil {
			log.Printf("Parse page: %v", err)
		}

	}
	fmt.Sprintln(content)
	return nil
}

// parsePage - parsing one page
func parsePage(urlToGet *url.URL, cfg *cfg.Config) string {

	// get content:
	content, err := getHTML(urlToGet.String(), cfg)
	if err != nil {
		log.Printf("Get HTML: %v, %v", urlToGet.String(), err)
	}

	// Parse images
	imgs, err := parseImages(urlToGet, content)
	if err != nil {
		log.Printf("Parse images: %v", err)
	}
	fmt.Sprintln(imgs)

	// parse links
	links, err := parseLinks(urlToGet, content)
	if err != nil {
		log.Printf("ParseLinks: %v", err)
	}
	fmt.Sprintln(links)

	return content
}

// parseImages - parse images from the page
func parseImages(urlToGet *url.URL, content string) ([]string, error) {
	var (
		err        error
		imgs       []string
		matches    [][]string
		findImages = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
		// `<img.*?src="(.*?)"`
		// `<img[^>]+\bsrc=["']([^"']+)["']`
	)

	// Retrieve all image URLs from string
	matches = findImages.FindAllStringSubmatch(content, -1)

	for _, val := range matches {
		var imgUrl *url.URL

		// Parse the image URL
		ur := strings.ReplaceAll(strings.ReplaceAll(val[1], "\n", ""), "\r", "")
		if imgUrl, err = url.Parse(ur); err != nil {
			log.Printf("Parse image url: %v, %v", val[1], err)
			continue
		}

		// ignore gif files
		if strings.HasPrefix(imgUrl.String(), "data:image/gif") {
			continue
		}

		// If the URL is absolute, add it to the slice
		// If the URL is relative, build an absolute URL
		if imgUrl.IsAbs() {
			imgs = append(imgs, imgUrl.String())
		} else {
			imgs = append(imgs, urlToGet.Scheme+"://"+urlToGet.Host+imgUrl.String())
		}

	}
	//TODO: УДАЛИТЬ
	for i := range imgs {
		fmt.Printf("img final: %s\n\n", imgs[i])
	}
	return imgs, err
}

func parseLinks(urlToGet *url.URL, content string) ([]string, error) {
	var (
		err       error
		links     []string = make([]string, 0)
		findLinks          = regexp.MustCompile("<a.*?href=\"(.*?)\"")
	)

	// Retrieve all anchor tag URLs from string
	matches := findLinks.FindAllStringSubmatch(content, -1)

	for _, val := range matches {
		var linkUrl *url.URL

		// Parse the anchr tag URL
		if linkUrl, err = url.Parse(val[1]); err != nil {
			log.Printf("Parse link url: %v", val[1])
			continue
		}

		// If the URL is absolute, add it to the slice
		// If the URL is relative, build an absolute URL
		if linkUrl.IsAbs() {
			links = append(links, linkUrl.String())
		} else {
			links = append(links, urlToGet.Scheme+"://"+urlToGet.Host+linkUrl.String())
		}
	}
	//TODO: УДАЛИТЬ
	// for i := range links {
	// 	fmt.Printf("links final: %s\n\n", links[i])
	// }

	return links, err
}
