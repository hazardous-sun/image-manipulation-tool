package image_editing

import (
	"image"
	"image/color"
	"log"
)

func FeatureExtractCountDominoDots(img image.Image) image.Image {
	found := false
	matrix := generateMatrix(img)
	count := 0

	// apply the grayscale filter
	tempImage := FilterGrayScale(img)

	// apply median blur
	tempImage = FilterMedianBlur(tempImage, 9)

	// apply sobel border detection
	tempImage = FilterSobelBorderDetection(tempImage, 255)

	// erode the image once
	tempImage = MathMorpDilation(tempImage)

	// erode the image twice
	tempImage = MathMorpErosion(tempImage)
	tempImage = MathMorpErosion(tempImage)

	// dilate the image thrice
	tempImage = MathMorpDilation(tempImage)
	tempImage = MathMorpDilation(tempImage)
	tempImage = MathMorpDilation(tempImage)

	resultImg := image.NewRGBA(img.Bounds())

	// iterate over all x and y, when a white spot is found, increase count
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			r, _, _, _ := tempImage.At(x, y).RGBA()
			if found == false && (r/256) >= 100 {
				if noWhiteNeighbours(x, y, matrix) {
					count++
					found = true
				}
			} else if found == true && (r/256) == 0 {
				found = false
			}
			if (r / 256) >= 100 {
				matrix[x][y] = 1
				pixel := color.NRGBA{
					R: 0,
					G: 0,
					B: 255,
					A: 255,
				}
				resultImg.Set(x, y, pixel)
			}
		}
	}

	log.Printf("Count: %d \n", count)
	return resultImg
}

func generateMatrix(img image.Image) [][]int {
	var matrix [][]int
	for x := 0; x < img.Bounds().Dx(); x++ {
		var row []int
		for y := 0; y < img.Bounds().Dy(); y++ {
			row = append(row, 0)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func noWhiteNeighbours(x, y int, matrix [][]int) bool {
	for innerX := -1; innerX < 1; innerX++ {
		for innerY := -1; innerY < 1; innerY++ {
			if matrix[x-innerX][y-innerY] == 1 {
				return false
			}
		}
	}
	return true
}

/*
Matrix with the size of the image
Set all values to 0
When a white pixel is found, check if the found flag is set, if it is not, check if there are white neighbour pixels
If there are no neighbour white pixels, increase count
*/
