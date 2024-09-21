package main

import (
	"fmt"
	"image"
	"image/color"
)

// Filters -------------------------------------------------------------------------------------------------------------

// -------- Grayscale

// Removes the color channels of an image and returns an image with only shades of gray.
//func invalidFilterGrayScale(img image.Image) image.Image {
//	grayImage := image.NewGray(img.Bounds())
//	draw.Draw(grayImage, grayImage.Bounds(), img, image.Point{}, draw.Src)
//	return grayImage
//}

// Removes the color channels of an image and returns an image with only shades of gray.
func filterGrayScale(img image.Image) image.Image {
	grayImage := image.NewGray(img.Bounds())
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			grayIntensity := (r + g + b) / 3
			grayImage.Set(x, y, color.Gray{
				Y: uint8(grayIntensity),
			})
		}
	}
	fmt.Println(grayImage.GrayAt(10, 10))
	return grayImage
}
