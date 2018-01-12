package main

import (
	"daily-desktop/scraper"
	"fmt"
)

// TODO figure out error handling

func main() {
	err := scraper.ScrapeSubredditForImage("earthporn")
	if err != nil {
		fmt.Println(err)
	}
}
