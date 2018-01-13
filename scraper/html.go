package scraper

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

/*

This module parses HTML from the subreddit for image tags

*/

/*
AppendDataURLToResult appends the url to an image on the subreddit
if the element contains a data-url to that image
*/
func AppendDataURLToResult(previousResult []string, token html.Token) []string {
	containsLink := false
	url := ""
	for _, attr := range token.Attr {
		if attr.Key == "class" {
			containsLink = strings.Index(attr.Val, " thing") != -1
			containsLink = containsLink && strings.Index(attr.Val, "link ") != -1
		}
		if attr.Key == "data-url" && (strings.Index(attr.Val, ".png") != -1 || strings.Index(attr.Val, ".jpg") != -1) {
			url = attr.Val
		}
	}
	if containsLink && url != "" {
		return append(previousResult, url)
	}
	return previousResult
}

/*
ScrapeImgUrlsFromHTML get urls of post pages from the HTML
of the subreddit
*/
func ScrapeImgUrlsFromHTML(responseBody *io.ReadCloser) ([]string, error) {
	tokenizer := html.NewTokenizer(*responseBody)
	result := make([]string, 0, 0)

	for {
		nextToken := tokenizer.Next()
		switch nextToken {
		case html.ErrorToken: // results slice at the end of the for loop
			return result, nil
		case html.StartTagToken, html.EndTagToken:
			token := tokenizer.Token()
			if token.Data == "div" {
				result = AppendDataURLToResult(result, token)
			}
		}
	}
}
