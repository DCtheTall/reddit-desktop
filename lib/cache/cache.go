package cache

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"reddit-desktop/lib/scraper"
	"sort"
	"time"
)

type timeSlice []time.Time

func (ts timeSlice) Len() int {
	return len(ts)
}

func (ts timeSlice) Less(i, j int) bool {
	return ts[i].Before(ts[j])
}

func (ts timeSlice) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

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
	if err != nil {
		return err
	}

	dirname = fmt.Sprintf("%s/data/", dirname)
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

func getPreviousImage(stepsBack int) (string, error) {
	dirname, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}

	files, err := ioutil.ReadDir(fmt.Sprintf("%s/data/", dirname))
	if err != nil {
		return "", err
	}
	if len(files) < stepsBack {
		return "", fmt.Errorf("Not enough images in your cache, run the scraper with the --save option to save images")
	}

	datesModified := make(timeSlice, len(files))
	for i, file := range files {
		datesModified[i] = file.ModTime()
	}

	sort.Sort(datesModified)
	previousFilesModTime := datesModified[len(datesModified)-stepsBack]
	var previousFile os.FileInfo
	for _, file := range files {
		if file.ModTime().Equal(previousFilesModTime) {
			previousFile = file
		}
	}

	return fmt.Sprintf("%s/data/%s", dirname, previousFile.Name()), nil
}

/*
GetPreviousImagePath gets the path previous image
from the cache
*/
func GetPreviousImagePath() (string, error) {
	return getPreviousImage(2)
}

/*
Pop pop the cache of the most recent image
*/
func Pop() error {
	filename, err := getPreviousImage(1)
	if err != nil {
		return err
	}

	err = os.RemoveAll(filename)
	if err != nil {
		return err
	}

	return nil
}
