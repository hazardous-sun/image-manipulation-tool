package main

import (
	"image"
	"math"
)

// Geometric transformations -------------------------------------------------------------------------------------------

func transformImage(img image.Image, matrix [][]int) image.Image {
	println("ENTREI EM transformImage()")
	transformedImage := image.NewRGBA(img.Bounds())
	// ---------------------------------------- Travel through img
	for x := 0; x < img.Bounds().Dx(); x++ { // -----------------+
		for y := 0; y < img.Bounds().Dy(); y++ { //--------------+
			pixelMatrix := []int{x, y, 1}
			tempValues := []int{0, 0, 0}
			// --------------------------------------------------------- Travel through pixelMatrix
			for row := 0; row < len(matrix); row++ { // ------------------------------------------+
				for column := 0; column < len(matrix[0]); column++ { // --------------------------+
					tempValues[row] += pixelMatrix[row] * matrix[row][column]
				}
			}
			transformedImage.Set(tempValues[0], tempValues[1], img.At(x, y))
		}
	}
	return transformedImage
}

// --------- Matrices

/*
Returns the matrix used for translating images:

[

	[1, 0, x],
	[0, 1, y],
	[0, 0, z],

]

Since we are dealing with images that only have two axis, we can set 'z' to 1.
*/
func getTranslationMatrix(x int, y int) [][]int {
	return [][]int{
		{1, 0, x},
		{0, 1, y},
		{0, 0, 1},
	}
}

/*
Returns the matrix used for scaling images:

[

	[x, 0, 0],
	[0, y, 0],
	[0, 0, z],

]

Since we are dealing with images that only have two axis, we can set 'z' to 1.
*/
func getScaleMatrix(x int, y int) [][]int {
	return [][]int{
		{x, 0, 0},
		{0, y, 0},
		{0, 0, 1},
	}
}

/*
Returns the matrix used to mirror images in the X axis:

[

	[0, 0, 0],
	[0, -1, 0],
	[0, 0, 1],

]
*/
func getMirrorHMatrix() [][]int {
	return [][]int{
		{0, 0, 0},
		{0, -1, 0},
		{0, 0, 1},
	}
}

/*
Returns the matrix used to mirror images in the Y axis:

[

	[0, 0, 0],
	[0, -1, 0],
	[0, 0, 1],

]
*/
func getMirrorVMatrix() [][]int {
	return [][]int{
		{-1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}

/*
Returns the matrix used for rotating images in the X axis:

[

	[1,       0,      0],
	[cos(α), -sen(α), 0],
	[sen(α),  cos(α), 0],

]
*/
func getRotationMatrix(x float64) [][]int {
	return [][]int{
		{1, 0, 0},
		{int(math.Cos(x)), -int(math.Sin(x)), 0},
		{int(math.Sin(x)), int(math.Cos(x)), 0},
	}
}
