package image_editing

import (
	"image"
	"image/color"
	"log"
)

const SpotValue = 10

func FeatureExtractCountDominoDots(img image.Image) image.Image {
	found := false
	alreadyIncremented := false
	rowCompleted := 0
	var positionsFound [][]int

	// apply the grayscale filter
	tempImage := FilterGrayScale(img)

	// apply median blur
	tempImage = FilterMedianBlur(tempImage, 9)

	// apply sobel border detection
	tempImage = FilterSobelBorderDetection(tempImage, 175)

	// erode the image once
	tempImage = MathMorpErosion(tempImage)

	resultImg := image.NewRGBA(img.Bounds())

	// iterate over all x and y, when a white spot is found, increase count
	for y := 1; y < img.Bounds().Dy()-5; y++ {
		for x := 1; x < img.Bounds().Dx()-5; x++ {
			r, _, _, _ := tempImage.At(x, y).RGBA()
			if found == false && (r/256) <= SpotValue {
				if alreadyIncremented == false {
					alreadyIncremented = true
					positionsFound = append(positionsFound, []int{x, y})
					rowCompleted++

					pixel := color.NRGBA{
						R: 0,
						G: 0,
						B: 255,
						A: 255,
					}
					resultImg.Set(x, y, pixel)
				}
				found = true
			} else if found == true && (r/256) > SpotValue {
				found = false
				if noSpotsNeighbours(x, y, tempImage) {
					alreadyIncremented = false
				}
			}
			if rowCompleted == 2 {
				rowCompleted = 0
				y += 5
				break
			}
		}
	}

	positionsFound = validatePositions(positionsFound)
	log.Printf("Count: %d \n", len(positionsFound))
	return resultImg
}

func noSpotsNeighbours(x, y int, img image.Image) bool {
	for innerX := -1; innerX < 1; innerX++ {
		for innerY := -1; innerY <= 1; innerY++ {
			r, _, _, _ := img.At(x-innerX, y-innerY).RGBA()
			newR := uint8(r / 256)
			if newR <= SpotValue {
				return false
			}
		}
	}
	return true
}

func validatePositions(positions [][]int) [][]int {
	removedValues := 0
	for i, value := range positions {
		if invalidPos(i, removedValues, value[0], value[1], positions) {
			positions = append(positions[:i-removedValues], positions[i+1-removedValues:]...)
			removedValues++
		}
	}
	return positions
}

func invalidPos(index, removedValues, x, y int, positions [][]int) bool {
	for i := -2; i <= 2; i++ {
		for j := -2; j <= 2; j++ {
			for innerIndex, v := range positions {
				if innerIndex-removedValues == index-removedValues || i == 0 && j == 0 {
					continue
				}
				if v[0] == x-i && v[1] == y-j {
					return true
				}
			}
		}
	}
	return false
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
