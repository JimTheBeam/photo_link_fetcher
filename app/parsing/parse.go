package parsing

import (
	"fmt"
	"log"
	"net/url"
	jsonparse "parse_photo_links/app/parsing/json_parse"
	"parse_photo_links/cfg"
	"strings"

	"golang.org/x/net/html"
)

var JSON = `{"url":["abc.com", "safasdf.com", "facebook.com","http://www.google.com/", "https://mail.ru/"]}`

// var JSON = `{"url":["httttps://github.com/"]}`

func ParseAll(cfg *cfg.Config) error {
	var (
		urls         jsonparse.IncomingJSON
		PagesContent []PageUrls
		content      PageUrls
	)

	// parse urls from incoming json and put them in urls
	if err := jsonparse.ParseJSON(JSON, &urls); err != nil {
		log.Printf("Parse json: %v", err)
		return errParseJson{}
	}

	// get page data:
	for _, oneUrl := range urls.Url {
		// Parse URL
		urlToGet, err := url.Parse(correctUrl(oneUrl))
		if err != nil {
			log.Printf("Parse url: %v\n", err)

			PagesContent = append(PagesContent,
				PageUrls{
					PageUrl:      urlToGet.String(),
					ErrorMessage: errUrlParse.Error(),
				},
			)
			continue
		}

		// get page content
		content, err = parsePage(urlToGet, cfg)
		if err != nil {
			PagesContent = append(PagesContent,
				PageUrls{
					PageUrl:      urlToGet.String(),
					ErrorMessage: err.Error(),
				},
			)
			continue
		}

		// append successfull parsing result in slice
		PagesContent = append(PagesContent, content)
	}

	// TODO: это печать она не нужна
	for i := range PagesContent {
		fmt.Printf("link: %s\nError: %v\nIMAGES: %v\n\n",
			PagesContent[i].PageUrl,
			PagesContent[i].ErrorMessage,
			PagesContent[i].Img)
	}
	return nil
}

// parsePage - parsing one page
func parsePage(urlToGet *url.URL, cfg *cfg.Config) (PageUrls, error) {

	var (
		links  []string
		images []string
	)

	// get content:
	content, err := getHTML(urlToGet.String(), cfg)
	if err != nil {
		log.Printf("Get HTML: %v, %v", urlToGet.String(), err)
		return PageUrls{}, err
	}

	// parse links
	parseLinks(&links, content, urlToGet)

	// parse images
	parseImg(&images, content, urlToGet)

	return PageUrls{
			PageUrl: urlToGet.String(),
			Img:     images,
			Link:    links,
		},
		nil
}

// parseLinks - parse links from the page
func parseLinks(links *[]string, content *html.Node, urlToGet *url.URL) []string {

	var err error

	if content.Type == html.ElementNode && content.Data == "a" {
		for _, a := range content.Attr {
			if a.Key == "href" {
				var linkUrl *url.URL

				// parse link url
				if linkUrl, err = url.Parse(a.Val); err != nil {
					log.Printf("Parse link url: %v", a.Val)
					continue
				}

				// If the URL is absolute, add it to the slice
				// If the URL is relative, build an absolute URL
				if linkUrl.IsAbs() {
					*links = append(*links, linkUrl.String())
				} else {
					*links = append(*links, urlToGet.Scheme+"://"+urlToGet.Host+linkUrl.String())
				}
			}
		}
	}

	// recursively find another links
	for c := content.FirstChild; c != nil; c = c.NextSibling {
		*links = parseLinks(links, c, urlToGet)
	}

	return *links
}

// parseImg - parse images from the page
func parseImg(links *[]string, content *html.Node, urlToGet *url.URL) []string {

	var err error

	if content.Type == html.ElementNode && content.Data == "img" {
		for _, img := range content.Attr {
			if img.Key == "src" {
				var imgUrl *url.URL

				// parse image url
				if imgUrl, err = url.Parse(img.Val); err != nil {
					log.Printf("Parse image url: %v", img.Val)
					continue
				}

				// ignore gif files
				if strings.HasPrefix(imgUrl.String(), "data:image/gif") {
					continue
				}

				// If the URL is absolute, add it to the slice
				// If the URL is relative, build an absolute URL
				if imgUrl.IsAbs() {
					*links = append(*links, imgUrl.String())
				} else {
					*links = append(*links, urlToGet.Scheme+"://"+urlToGet.Host+imgUrl.String())
				}
			}
		}
	}

	// recursively find another img links
	for c := content.FirstChild; c != nil; c = c.NextSibling {
		*links = parseImg(links, c, urlToGet)
	}

	return *links
}
