package image_handling_test

import (
	"image-manipulation-tool/image-handling"
	"math"
	"testing"
)

func TestGetTranslationMatrix(t *testing.T) {
	x := 3.0
	y := 5.0
	receivedMatrix := image_handling.GetTranslationMatrix(x, y)
	expectedMatrix := [][]float64{
		{1, 0, x},
		{0, 1, y},
		{0, 0, 1},
	}
	for i := range receivedMatrix {
		for j := range receivedMatrix[i] {
			if receivedMatrix[i][j] != expectedMatrix[i][j] {
				t.Errorf(RError()+" expected translation receivedMatrix element at [%d][%d] to be %f, but got %f", i, j, expectedMatrix[i][j], receivedMatrix[i][j])
			}
		}
	}
}

func TestGetResizeMatrix(t *testing.T) {
	x := 3.0
	y := 5.0
	receivedMatrix := image_handling.GetResizeMatrix(x, y)
	expectedMatrix := [][]float64{
		{1 / x, 0, 0},
		{0, 1 / y, 0},
		{0, 0, 1},
	}
	for i := range receivedMatrix {
		for j := range receivedMatrix[i] {
			if receivedMatrix[i][j] != expectedMatrix[i][j] {
				t.Errorf(RError()+" expected translation receivedMatrix element at [%d][%d] to be %f, but got %f", i, j, expectedMatrix[i][j], receivedMatrix[i][j])
			}
		}
	}
}

func TestGetMirrorHMatrix(t *testing.T) {
	receivedMatrix := image_handling.GetMirrorHMatrix()
	expectedMatrix := [][]float64{
		{-1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	for i := range receivedMatrix {
		for j := range receivedMatrix[i] {
			if receivedMatrix[i][j] != expectedMatrix[i][j] {
				t.Errorf(RError()+" expected translation receivedMatrix element at [%d][%d] to be %f, but got %f", i, j, expectedMatrix[i][j], receivedMatrix[i][j])
			}
		}
	}
}

func TestGetMirrorVMatrix(t *testing.T) {
	receivedMatrix := image_handling.GetMirrorVMatrix()
	expectedMatrix := [][]float64{
		{1, 0, 0},
		{0, -1, 0},
		{0, 0, 1},
	}
	for i := range receivedMatrix {
		for j := range receivedMatrix[i] {
			if receivedMatrix[i][j] != expectedMatrix[i][j] {
				t.Errorf(RError()+" expected translation receivedMatrix element at [%d][%d] to be %f, but got %f", i, j, expectedMatrix[i][j], receivedMatrix[i][j])
			}
		}
	}
}

func TestGetRotationMatrix(t *testing.T) {
	x := 3.0
	receivedMatrix := image_handling.GetRotationMatrix(x)
	expectedMatrix := [][]float64{
		{math.Cos(x), math.Sin(x), 0},
		{-math.Sin(x), math.Cos(x), 0},
		{0, 0, 1},
	}
	for i := range receivedMatrix {
		for j := range receivedMatrix[i] {
			if receivedMatrix[i][j] != expectedMatrix[i][j] {
				t.Errorf(RError()+" expected translation receivedMatrix element at [%d][%d] to be %f, but got %f", i, j, expectedMatrix[i][j], receivedMatrix[i][j])
			}
		}
	}
}

func RError() string {
	return "\033[31merror:\033[0m"
}
