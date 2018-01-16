package scraper

import "image"

/*
ScrapedImage from the web
with a recorded name
*/
type ScrapedImage struct {
	img       *image.Image
	name      string
	extension string
}

/*
GetImage gets Image object stored in struct
*/
func (s *ScrapedImage) GetImage() *image.Image {
	return s.img
}

/*
GetName gets name from struct
*/
func (s *ScrapedImage) GetName() string {
	return s.name
}

/*
GetExtension of the image stored in the struct
*/
func (s *ScrapedImage) GetExtension() string {
	return s.extension
}
