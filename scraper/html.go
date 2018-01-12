package scraper

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

/*

This module parses HTML from the subreddit for image tags

*/

/*
AppendURLIfPostThumbnail returns list with new
url if the anchor tag links to a post page
*/
func AppendURLIfPostThumbnail(previousResult []string, token html.Token) []string {
	isThumbnail := false
	url := ""
	for _, attr := range token.Attr {
		if attr.Key == "class" {
			isThumbnail = strings.Index(attr.Val, "thumbnail") != -1
		}
		if attr.Key == "href" {
			url = fmt.Sprintf("https://www.reddit.com%s", attr.Val)
		}
	}
	if isThumbnail && len(previousResult) < 20 {
		return append(previousResult, url)
	}
	return previousResult
}

/*
ScrapePostUrlsFromHTML get urls of post pages from the HTML
of the subreddit
*/
func ScrapePostUrlsFromHTML(responseBody io.ReadCloser) ([]string, error) {
	tokenizer := html.NewTokenizer(responseBody)
	result := make([]string, 0, 0)
	openingAnchorToken := html.Token{}
	anchorTagOpen := false
	hasImgChild := false
	for {
		nextToken := tokenizer.Next()
		switch nextToken {
		case html.ErrorToken: // results slice at the end of the for loop
			return result, nil
		case html.StartTagToken, html.EndTagToken:
			token := tokenizer.Token()
			switch {
			case token.Data == "a" && strings.Index(token.String(), "<a") != -1:
				anchorTagOpen = true
				openingAnchorToken = token
			case token.Data == "a" && strings.Index(token.String(), "a>") != -1:
				anchorTagOpen = false
				if hasImgChild {
					result = AppendURLIfPostThumbnail(result, openingAnchorToken)
				}
			case token.Data == "img" && anchorTagOpen:
				hasImgChild = true
			}
		}
	}
}
