package scraper

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

/*

This module parses HTML from the subreddit for image tags

So far: gotten tagnames
get the elements with class names that link to images

*/

// TODO
func GetPostURL() string {
	return ""
}

func GetPostUrlsFromHTML(responseBody io.ReadCloser) ([]string, error) {
	tokenizer := html.NewTokenizer(responseBody)
	result := make([]string, 0, 0)
	for {
		token := tokenizer.Next()
		switch token {
		case html.ErrorToken:
			return result, nil
		case html.StartTagToken, html.EndTagToken:
			tagnameBytes, _ := tokenizer.TagName()
			tagname := string(tagnameBytes)
			if tagname == "a" {
				result = append(result, GetPostURL())
			}
			fmt.Println(tagname)
		}
	}
}
