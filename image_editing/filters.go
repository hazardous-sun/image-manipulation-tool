package image_editing

import (
	"image"
	"image/color"
)

// Filters -------------------------------------------------------------------------------------------------------------

// FilterGrayScale :
// Removes the color channels of an image and returns an image with only shades of gray.
func FilterGrayScale(img image.Image) image.Image {
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

func FilterContrast(img image.Image, contrast float64) image.Image {
	contrastImage := image.NewRGBA(img.Bounds())
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			r, g, b, a := img.At(x, y).RGBA()
			newR := float64(r) * contrast
			newG := float64(g) * contrast
			newB := float64(b) * contrast
			pixel := color.RGBA{
				R: uint8(newR / 256),
				G: uint8(newG / 256),
				B: uint8(newB / 256),
				A: uint8(a),
			}
			contrastImage.Set(x, y, pixel)
		}
	}
	return contrastImage
}
