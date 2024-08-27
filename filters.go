package main

import (
	"image"
	"image/draw"
)

// Filters -------------------------------------------------------------------------------------------------------------

// -------- Grayscale
/*
Applies the grayscale filter to an image.
*/
func filterGrayScale(img image.Image) image.Image {
	grayImage := image.NewGray(img.Bounds())
	draw.Draw(grayImage, grayImage.Bounds(), img, image.Point{}, draw.Src)
	return grayImage
}

//func oldFilterGrayScale(img image.Image) image.Image {
//	grayImage := image.NewGray(img.Bounds())
//	for x := 0; x < img.Bounds().Dx(); x++ {
//		for y := 0; y < img.Bounds().Dy(); y++ {
//			r, g, b, _ := img.At(x, y).RGBA()
//			grayIntensity := (r + g + b) / 3
//			grayImage.Set(x, y, color2.Gray{
//				Y: uint8(grayIntensity),
//			})
//		}
//	}
//	fmt.Println(grayImage.GrayAt(10, 10))
//	return grayImage
//}
