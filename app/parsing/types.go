package parsing

type PageUrls struct {
	PageUrl      string   `json:"url"`
	ErrorMessage string   `json:"errorMessage"`
	Img          []string `json:"images"`
	Link         []string `json:"links"`
}
