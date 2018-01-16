package main

import (
	"daily-desktop/desktopimage"
	"daily-desktop/scraper"
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
	desktopimage.GetImageForDesktop(images)

	// Apple Script for setting bg:
	// osascript -e 'tell application "Finder" to set desktop picture to "/Users/you/Pictures/Some Picture.jpg" as POSIX file'

	// out, _ := exec.Command("echo", "Hello World!").Output()
}
