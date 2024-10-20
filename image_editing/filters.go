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
			tempR := float64(r) / 256 * contrast
			if tempR > 255 {
				if tempR > 279 {
					tempR = 255
				} else {
					tempR = 255 - 24
				}
			} else if tempR < 0 {
				tempR = 0
			} else {
				tempR -= 72
				if tempR < 0 {
					tempR = 0
				}
			}

			tempG := float64(g) / 256 * contrast
			if tempG > 255 {
				if tempG > 279 {
					tempG = 255
				} else {
					tempG = 255 - 24
				}
			} else if tempG < 0 {
				tempG = 0
			} else {
				tempG -= 72
				if tempG < 0 {
					tempG = 0
				}
			}

			tempB := float64(b) / 256 * contrast
			if tempB > 255 {
				if tempB > 279 {
					tempB = 255
				} else {
					tempB = 255 - 24
				}
			} else if tempB < 0 {
				tempB = 0
			} else {
				tempB -= 72
				if tempB < 0 {
					tempB = 0
				}
			}

			newR := uint8(tempR)
			newG := uint8(tempG)
			newB := uint8(tempB)

			pixel := color.RGBA{
				R: newR,
				G: newG,
				B: newB,
				A: uint8(a),
			}
			contrastImage.Set(x, y, pixel)
		}
	}
	return contrastImage
}
