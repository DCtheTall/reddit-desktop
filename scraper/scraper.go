package scraper

import "fmt"

/*

Scraper gets the image data from Reddit

TODO parse comment pages for the link to the image,
start by writing a function to request the

*/

/*
ScrapeSubredditForImage will return image data when done
*/
func ScrapeSubredditForImage(subreddit string) []error {
	url := fmt.Sprintf("https://www.reddit.com/r/%s", subreddit)
	subredditBody, err := GetRedditPage(url)
	if err != nil {
		return []error{err}
	}
	defer (*subredditBody).Close()

	urls, err := ScrapeImgUrlsFromHTML(subredditBody)
	if err != nil {
		return []error{err}
	}
	fmt.Println(urls)
	return nil
}
