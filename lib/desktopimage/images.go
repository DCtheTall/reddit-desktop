package desktopimage

import (
	"fmt"
	"math/rand"
	"reddit-desktop/lib/scraper"
	"time"
)

/*
GetImageForDesktop filters slice of image pointers and
returns a random image suitable for the desktop
*/
func GetImageForDesktop(images []*scraper.ScrapedImage) *scraper.ScrapedImage {
	validImages := make([]*scraper.ScrapedImage, 0, 0)
	for _, scrapedImg := range images {
		img := *scrapedImg.GetImage()
		bounds := img.Bounds()
		if bounds.Max.X > 1000 {
			validImages = append(validImages, scrapedImg)
		}
	}
	fmt.Println(fmt.Sprintf("\nFiltered %d possible images out of %d images", len(validImages), len(images)))
	rand.Seed(time.Now().UnixNano())
	return validImages[rand.Intn(len(validImages))]
}
