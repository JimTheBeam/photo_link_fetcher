### Parse photo and links from web sites

for example:

host/api/v1/fetch

takes json: {"url":["abc.com","facebook.com","google.com"]}

return json with all photos and links from every site in json array:

{"error": "an error",
    "data": [
        {
            "url": "url of the page",
            "errorMessage": "error message for the page",
            "images": ["array of images urls from the page",],
            "links": ["array of links urls from the page",]
        },
}