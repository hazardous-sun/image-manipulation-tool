package image_editing

import (
	"image"
	"math"
)

// Geometric transformations -------------------------------------------------------------------------------------------

func TransformImage(img image.Image, matrix [][]float64) image.Image {
	transformedImage := image.NewRGBA(img.Bounds())
	// ---------------------------------------- Travel through img
	for x := 0; x < img.Bounds().Dx(); x++ { // -----------------+
		for y := 0; y < img.Bounds().Dy(); y++ { //--------------+
			applyChange(img, transformedImage, matrix, x, y)
		}
	}
	return transformedImage
}

func applyChange(img image.Image, transformedImage *image.RGBA, matrix [][]float64, x int, y int) {
	halfX := img.Bounds().Dx() / 2
	halfY := img.Bounds().Dy() / 2
	tmpX := float64(x - halfX)
	tmpY := float64(y - halfY)
	newX := tmpX*matrix[0][0] + tmpY*matrix[0][1] + 1*matrix[0][2]
	newY := tmpX*matrix[1][0] + tmpY*matrix[1][1] + 1*matrix[1][2]
	newX += float64(halfX)
	newY += float64(halfY)

	if newX < float64(img.Bounds().Dx()) && newY < float64(img.Bounds().Dy()) && newX >= 0 && newY >= 0 {
		transformedImage.Set(x, y, img.At(int(newX), int(newY)))
	}
}

// --------- Matrices

// GetTranslationMatrix :
// Returns the matrix used for translating images:
// [
//
//	[1, 0, x],
//	[0, 1, y],
//	[0, 0, z],
//
// ]
// Since we are dealing with images that only have two axis, we can set 'z' to 1.
func GetTranslationMatrix(x float64, y float64) [][]float64 {
	return [][]float64{
		{1, 0, x},
		{0, 1, y},
		{0, 0, 1},
	}
}

// GetResizeMatrix :
// Returns the matrix used for scaling images:
//
// [
//
//	[1/x,   0, 0],
//	[  0, 1/y, 0],
//	[  0,   0, z],
//
// ]
//
// Since we are dealing with images that only have two axis, we can set 'z' to 1.
func GetResizeMatrix(x float64, y float64) [][]float64 {
	return [][]float64{
		{1 / x, 0, 0},
		{0, 1 / y, 0},
		{0, 0, 1},
	}
}

// GetMirrorHMatrix :
// Returns the matrix used to mirror images in the X axis:
//
// [
//
//	[-1, 0, 0],
//	[0, 1, 0],
//	[0, 0, 1],
//
// ]
func GetMirrorHMatrix() [][]float64 {
	return [][]float64{
		{-1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}

// GetMirrorVMatrix :
// Returns the matrix used to mirror images in the Y axis:
//
// [
//
//	[1, 0, 0],
//	[0, -1, 0],
//	[0, 0, 1],
//
// ]
func GetMirrorVMatrix() [][]float64 {
	return [][]float64{
		{1, 0, 0},
		{0, -1, 0},
		{0, 0, 1},
	}
}

// GetRotationMatrix :
// Returns the matrix used for rotating images in the X axis:
//
// [
//
//	[ cos(α), sin(α), 0],
//	[-sin(α), cos(α), 0],
//	[      0,      0, 1],
//
// ]
func GetRotationMatrix(x float64) [][]float64 {
	return [][]float64{
		{math.Cos(x), math.Sin(x), 0},
		{-math.Sin(x), math.Cos(x), 0},
		{0, 0, 1},
	}
}
