package cache

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"reddit-desktop/lib/scraper"
)

/*
Save saves the image in the data/ directory located
in the directory the app runs in
*/
func Save(image *scraper.ScrapedImage) (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	cwd := filepath.Dir(ex)
	filename := fmt.Sprintf("%s/data/%s", cwd, image.GetName())
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

/*
EmptyCache of all stored images

based on:
https://stackoverflow.com/questions/33450980/golang-remove-all-contents-of-a-directory
*/
func EmptyCache() error {
	dirname, err := filepath.Abs(filepath.Dir(os.Args[0]))
	dirname = fmt.Sprintf("%s/data/", dirname)
	if err != nil {
		return err
	}
	dir, err := os.Open(dirname)
	if err != nil {
		return err
	}
	defer dir.Close()
	names, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dirname, name))
		if err != nil {
			return err
		}
	}
	return nil
}
