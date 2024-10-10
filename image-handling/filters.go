package image_handling

import (
	"image"
	"image/color"
)

// Filters -------------------------------------------------------------------------------------------------------------

// -------- Grayscale

// Removes the color channels of an image and returns an image with only shades of gray.

// Removes the color channels of an image and returns an image with only shades of gray.
func filterGrayScale(img image.Image) image.Image {
	grayImage := image.NewGray(img.Bounds())
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			lum := 0.2125*float64(r) + 0.7154*float64(g) + 0.0721*float64(b)
			pixel := color.Gray{Y: uint8(lum / 256)}
			grayImage.Set(x, y, pixel)
		}
	}
	return grayImage
}
