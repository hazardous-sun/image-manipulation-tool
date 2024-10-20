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

func FilterContrast(img image.Image, factor float64) image.Image {
	if factor == 0 {
		return img
	}

	contrastImage := image.NewRGBA(img.Bounds())
	contrast := factor / 50
	if contrast < 0 {
		contrast *= -1
	}
	removedValue := 16896 * (factor * contrast / 100)
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			r, g, b, a := img.At(x, y).RGBA()

			newR := uint8(getNewChannelVal(r, contrast, removedValue))
			newG := uint8(getNewChannelVal(g, contrast, removedValue))
			newB := uint8(getNewChannelVal(b, contrast, removedValue))

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

func getNewChannelVal(x uint32, contrast float64, removedValue float64) float64 {
	temp := (float64(x)*contrast - removedValue) / 256
	if temp > 255 {
		temp = 255
	} else if temp < 0 {
		temp = 0
	}
	return temp
}
