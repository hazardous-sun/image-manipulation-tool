package main

import "image"

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

func getTranslationMatrix(x int, y int) [][]int {
	matrix := [][]int{
		{1, 0, x},
		{0, 1, y},
		{0, 0, 1},
	}
	return matrix
}
