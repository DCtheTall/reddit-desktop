package desktopimage

import (
	"image"
	"math/rand"
)

/*
GetImageForDesktop filters slice of image pointers and
returns a random image suitable for the desktop

TODO: let the minimum width (or height) be a cmd line argument
*/
func GetImageForDesktop(images []*image.Image) (img *image.Image) {
	validImages := make([]*image.Image, 0, 0)
	for _, img := range images {
		bounds := (*img).Bounds()
		if bounds.Max.X > bounds.Max.Y && bounds.Max.X > 1200 {
			validImages = append(validImages, img)
		}
	}
	index := rand.Intn(len(validImages))
	return validImages[index]
}
