package main

import (
	"daily-desktop/desktopimage"
	"daily-desktop/scraper"
	"fmt"
	"log"
	"os"
	"time"
)

// TODO figure out error handling

func main() {
	images, errs := scraper.ScrapeSubredditForImages("earthporn")
	if errs != nil && len(errs) > 0 {
		for _, err := range errs {
			log.Fatal(err)
		}
	}
	img := desktopimage.GetImageForDesktop(images)

	filename, err := desktopimage.WriteImageToFile(img)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		time.Sleep(5e3 * time.Millisecond)
		err := os.Remove(filename)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("Deleted %s", img.GetName()))
	}()

	err = desktopimage.SetDesktopBackground(filename)
	if err != nil {
		log.Fatal(err)
	}
}
