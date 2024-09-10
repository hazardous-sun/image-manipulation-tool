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
			applyChange(img, transformedImage, matrix, x, y)
		}
	}
	return transformedImage
}

func applyChange(img image.Image, transformedImage *image.RGBA, matrix [][]int, x int, y int) {
	halfX := img.Bounds().Dx() / 2
	halfY := img.Bounds().Dy() / 2
	tmpX := x - halfX
	tmpY := y - halfY
	newX := tmpX*matrix[0][0] + tmpY*matrix[0][1] + 1*matrix[0][2]
	newY := tmpX*matrix[1][0] + tmpY*matrix[1][1] + 1*matrix[1][2]
	newX += halfX
	newY += halfY

	if newX < img.Bounds().Dx() && newY < img.Bounds().Dy() && newX >= 0 && newY >= 0 {
		transformedImage.Set(x, y, img.At(newX, newY))
	}
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
		{-1, 0, 0},
		{0, 1, 0},
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
		{1, 0, 0},
		{0, -1, 0},
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
