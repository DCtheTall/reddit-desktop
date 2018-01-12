package scraper

import "fmt"

/*
ScrapeSubredditForImage will return image data when done
*/
func ScrapeSubredditForImage(subreddit string) error {
	resBody, err := GetSubredditPage(subreddit)
	if err != nil {
		return err
	}
	defer resBody.Close()
	urls, _ := ScrapePostUrlsFromHTML(resBody)
	fmt.Println(urls)
	return nil
}
