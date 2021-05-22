package parsing

import "strings"

// correctUrl - check if url is correct type: "http://www.google.com"
func correctUrl(url string) string {
	url = strings.TrimLeft(strings.TrimLeft(url, "http://"), "https://")
	return "https://www." + strings.TrimRight(strings.TrimLeft(url, "www."), "/")
}
