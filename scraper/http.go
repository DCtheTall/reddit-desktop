package scraper

import (
	"bytes"
	"image"
	_ "image/jpeg" // need these for decoding images
	_ "image/png"
	"io"
	"net/http"
)

/*

This module contains the HTTP request for
getting a subreddit view's HTML

*/

/*
GetImgFromURL fetches an image at the given url and
returns a pointer to an image object with the data
*/
func GetImgFromURL(url string) (*image.Image, error) {
	var client http.Client

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("user-agent", "daily-desktop-bot")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}

	return &img, nil
}

/*
GetRedditPage gets the HTML for the Reddit
page at the given URL
*/
func GetRedditPage(url string) (*io.ReadCloser, error) {
	var client http.Client
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("user-agent", "daily-desktop-bot")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return &resp.Body, nil
}

/*
CovertResponseBodyToString useful for looking at HTML
in the console
*/
func CovertResponseBodyToString(responseBody io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(responseBody)
	return buf.String()
}
