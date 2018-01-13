package scraper

import (
	"errors"
	"fmt"
	"image"
)

/*

Scraper gets the image data from Reddit

TODO parse comment pages for the link to the image,
start by writing a function to request the

*/

/*
ScrapeSubredditForImage will return image data when done
*/
func ScrapeSubredditForImages(subreddit string) ([]*image.Image, []error) {
	url := fmt.Sprintf("https://www.reddit.com/r/%s", subreddit)
	subredditBody, err := GetRedditPage(url)
	if err != nil {
		return nil, []error{err}
	}
	defer (*subredditBody).Close()

	urls, err := ScrapeImgUrlsFromHTML(subredditBody)
	if err != nil {
		return nil, []error{err}
	}

	if len(urls) > 0 {
		_, err = GetImgFromURL(urls[0])
		if err != nil {
			return nil, []error{err}
		}
	} else {
		return nil, []error{errors.New("No images available")}
	}
	return nil, nil
}
