package main

import (
	"daily-desktop/scraper"
	"fmt"
	"log"
)

// TODO figure out error handling

func main() {
	images, errs := scraper.ScrapeSubredditForImages("earthporn")
	if errs != nil && len(errs) > 0 {
		for _, err := range errs {
			log.Fatal(err)
		}
	}

	// Temporary, images will be handed in module that sets background
	for _, img := range images {
		bounds := (*img).Bounds()
		fmt.Println(bounds.Max.X, bounds.Max.Y)
	}
}
