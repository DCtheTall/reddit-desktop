package main

import (
	// "daily-desktop/scraper"
	// "fmt"
	"daily-desktop/scraper"
	"fmt"
)

// TODO figure out error handling

func main() {
	errs := scraper.ScrapeSubredditForImage("earthporn")
	if errs != nil && len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
	}
}
