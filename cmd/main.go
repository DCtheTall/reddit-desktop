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
	}

	if options[args.Undo] {
		filename, err := cache.GetPreviousImagePath()
		if err != nil {
			log.Fatal(err)
		}
		desktopimage.SetDesktopBackground(filename)
		if err := cache.Pop(); err != nil {
			log.Fatal(err)
		}
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

	filename, err := cache.Save(img)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}

		fmt.Println()
		time.Sleep(2e3 * time.Millisecond)

		if !options[args.Cache] {
			if err := os.Remove(filename); err != nil {
				log.Fatal(err)
			}
			fmt.Println("Deleted ", filename)
		}
	}()

	if err := desktopimage.SetDesktopBackground(filename); err != nil {
		log.Panic(err)
	}
	fmt.Println("Done")
}
