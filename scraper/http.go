package scraper

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

/*

This module contains the HTTP request for
getting a subreddit view's HTML

*/

/*
GetSubredditPage makes GET request to r/<subreddit>
returns open response body from GET request
*/
func GetSubredditPage(subreddit string) (io.ReadCloser, error) {
	var client http.Client
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.reddit.com/r/%s", subreddit), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("user-agent", "daily-desktop-bot")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
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
