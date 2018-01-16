package desktopimage

import (
	"daily-desktop/scraper"
	"fmt"
	"image/jpeg"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

/*
GetImageForDesktop filters slice of image pointers and
returns a random image suitable for the desktop

TODO: let the minimum width (or height) be a cmd line argument
*/
func GetImageForDesktop(images []*scraper.ScrapedImage) *scraper.ScrapedImage {
	validImages := make([]*scraper.ScrapedImage, 0, 0)
	for _, scrapedImg := range images {
		img := scrapedImg.GetImage()
		bounds := (*img).Bounds()
		if bounds.Max.X > 1000 {
			validImages = append(validImages, scrapedImg)
		}
	}
	fmt.Println(fmt.Sprintf("\nFiltered %d possible images out of %d images", len(validImages), len(images)))
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(validImages))
	fmt.Println(index)
	return validImages[index]
}

/*
WriteImageToFile write a scraped image to a file on the OS
*/
func WriteImageToFile(image *scraper.ScrapedImage) (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	cwd := filepath.Dir(ex)
	filename := fmt.Sprintf("%s/%s", cwd, image.GetName())
	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}

	defer func() {
		fmt.Println(fmt.Sprintf("Successfully saved %s", filename))
		file.Close()
	}()

	switch image.GetExtension() {
	case "jpeg":
		err = jpeg.Encode(file, *image.GetImage(), &jpeg.Options{jpeg.DefaultQuality})
	case "png":
		err = png.Encode(file, *image.GetImage())
	}
	if err != nil {
		return "", err
	}

	return filename, nil
}
