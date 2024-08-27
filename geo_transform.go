package main

import (
	"image"
	"math"
)

// Geometric transformations -------------------------------------------------------------------------------------------

func transformImage(img image.Image, matrix [][]int) image.Image {
	transformedImage := image.NewRGBA(img.Bounds())

	/*
						[
							[v1, v2, v3]
		[Xf, Yf, 1] =		[v1, v2, v3]   *   [X1, Y1, 1]
							[v1, v2, v3]
						]
	*/

	return transformedImage
}

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
