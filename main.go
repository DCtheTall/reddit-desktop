package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/DCtheTall/reddit-desktop/desktopimage"
	"github.com/DCtheTall/reddit-desktop/scraper"
)

// TODO figure out error handling

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		log.Fatal(errors.New("You must supply at least one subreddit to choose from"))
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(argsWithoutProg))
	subreddit := argsWithoutProg[index]

	images, errs := scraper.ScrapeSubredditForImages(subreddit)
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
		err := recover()
		log.Println(err)
		fmt.Println()
		time.Sleep(5e3 * time.Millisecond)
		err = os.Remove(filename)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("Deleted %s", img.GetName()))
	}()

	err = desktopimage.SetDesktopBackground(filename)
	if err != nil {
		log.Panic(err)
	}
}
