package server

import "parse_photo_links/app/parsing"

type ResponseJson struct {
	Error string             `json:"error"`
	Data  []parsing.PageUrls `json:"data"`
}
