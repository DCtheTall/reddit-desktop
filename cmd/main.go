package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reddit-desktop/lib/args"
	"reddit-desktop/lib/cache"
	"reddit-desktop/lib/desktopimage"
	"reddit-desktop/lib/scraper"
	"strings"
	"time"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		log.Fatal(errors.New("You must supply at least one subreddit to choose from"))
	}

	subreddits, options, err := args.ParseArgs(argsWithoutProg)
	if err != nil {
		log.Fatal(err)
	}

	if options[args.EmptyCache] {
		err = cache.EmptyCache()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	fmt.Println("Scraping subreddits: ", strings.Join(subreddits, ", "))
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(subreddits))
	subreddit := subreddits[index]

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
		if err != nil {
			log.Println(err)
		}

		fmt.Println()
		time.Sleep(2e3 * time.Millisecond)

		if !options[args.Cache] {
			err = os.Remove(filename)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Deleted ", filename)
		}
	}()

	err = desktopimage.SetDesktopBackground(filename)
	if err != nil {
		log.Panic(err)
	}
}
