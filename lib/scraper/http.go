package scraper

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg" // need these for decoding images
	_ "image/png"
	"io"
	"net/http"
	"strings"
	"sync"
)

/*

This module contains the HTTP request for
getting a subreddit view's HTML

*/

/*
CovertResponseBodyToString useful for looking at HTML
in the console
*/
func covertResponseBodyToString(responseBody *io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(*responseBody)
	return buf.String()
}

/*
GetRedditPage gets the HTML for the Reddit
page at the given URL

returned response body is still open
*/
func GetRedditPage(url string) (*io.ReadCloser, error) {
	fmt.Println("Requesting HTML from ", url)
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
	fmt.Println("\nRequest for ", url, " successful")
	return &resp.Body, nil
}

/*
getImgFromURL fetches an image at the given url and
returns a pointer to an image object with the data
*/
func getImgFromURL(url string) (*image.Image, string, error) {
	var client http.Client

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", err
	}

	req.Header.Set("user-agent", "daily-desktop-bot")
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	img, ext, err := image.Decode(resp.Body)
	if err != nil {
		return nil, "", err
	}

	return &img, ext, nil
}

/*
imageCollector uses sync.Mutex to allow safe
concurrent access to the error and Image data
arrays
*/
type imageCollector struct {
	access sync.Mutex
	data   []*ScrapedImage
	errors []error
}

/*
GetImagesFromScrapedURLs gets images from the scraped
source urls from the subreddit page
*/
func GetImagesFromScrapedURLs(urls []string) ([]*ScrapedImage, []error) {
	var wg sync.WaitGroup
	images := imageCollector{}
	for _, url := range urls {
		url := url
		splitURL := strings.Split(url, "/")
		name := splitURL[len(splitURL)-1]
		wg.Add(1)
		go func() {
			defer wg.Done()
			images.access.Lock()
			defer images.access.Unlock()
			if img, extension, err := getImgFromURL(url); err != nil {
				images.errors = append(images.errors, err)
			} else {
				fmt.Println("Successfully retrieved image from ", url)
				images.data = append(images.data, &ScrapedImage{
					img:       img,
					name:      name,
					extension: extension,
				})
			}
		}()
	}
	wg.Wait()
	return images.data, images.errors
}
