package scraper

import (
	"errors"
	"fmt"
)

/*

Scraper gets the image data from Reddit

*/

/*
ScrapeSubredditForImages scrapes image data from links
on the subreddit page
*/
func ScrapeSubredditForImages(subreddit string) ([]*ScrapedImage, []error) {
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
	for _, url := range urls {
		fmt.Println("Scraped image url: ", url)
	}

	if len(urls) == 0 {
		return nil, []error{errors.New("No images available")}
	}
	imgs, errs := GetImagesFromScrapedURLs(urls)
	if len(errs) > 0 {
		return nil, errs
	}
	return imgs, nil
}
