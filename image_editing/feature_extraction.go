package image_editing

import (
	"image"
	"image/color"
	"log"
)

func FeatureExtractCountDominoDots(img image.Image) image.Image {
	found := false
	alreadyIncremented := false
	rowCompleted := 0
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

	// new ------------------------------------------------------
	// median blur
	tempImage = FilterMedianBlur(tempImage, 9)
	// dilation
	tempImage = MathMorpDilation(tempImage)
	// new ------------------------------------------------------

	//tempImage = MathMorpErosion(tempImage)
	//
	//// dilate the image thrice
	//tempImage = MathMorpDilation(tempImage)
	//tempImage = MathMorpDilation(tempImage)
	//tempImage = MathMorpDilation(tempImage)
	//
	resultImg := image.NewRGBA(img.Bounds())

	// iterate over all x and y, when a white spot is found, increase count
	for y := 0; y < img.Bounds().Dx(); y++ {
		for x := 0; x < img.Bounds().Dy(); x++ {
			r, _, _, _ := tempImage.At(x, y).RGBA()
			if found == false && (r/256) >= 100 {
				if alreadyIncremented == false {
					alreadyIncremented = true
					count++
					rowCompleted++
				}
				found = true
			} else if found == true && (r/256) == 0 {
				found = false
				if noWhiteNeighboursUpdated(x, y, tempImage) {
					alreadyIncremented = false
				}
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
			if rowCompleted == 2 {
				rowCompleted = 0
				y += 5
				break
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
	for innerX := -2; innerX < 2; innerX++ {
		for innerY := -2; innerY < 2; innerY++ {
			if matrix[x-innerX][y-innerY] == 1 {
				return false
			}
		}
	}
	return true
}

func noWhiteNeighboursUpdated(x, y int, img image.Image) bool {
	for innerX := -1; innerX < 1; innerX++ {
		for innerY := -1; innerY < 1; innerY++ {
			r, _, _, _ := img.At(x-innerX, y-innerY).RGBA()
			newR := uint8(r / 256)
			if newR >= 255 {
				return false
			}
		}
	}
	return true
}

/*
grayscale
sobel 255
dilation
erosion
erosion
dilation
dilation
dilation

iterate over all x and y, when a white spot is found, increase count


x = 52
y = 24

52 - -1 = 53
52 - 0 = 52
52 - 1 = 51


Matrix with the size of the image
Set all values to 0
When a white pixel is found, check if the found flag is set, if it is not, check if there are white neighbour pixels
If there are no neighbour white pixels, increase count
*/
