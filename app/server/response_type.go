package server

import "parse_photo_links/app/parsing"

type ResponseJson struct {
	ErrorCode    int                `json:"errorCode"`
	ErrorMessage string             `json:"errorMessage"`
	Result       []parsing.PageUrls `json:"data"`
}
