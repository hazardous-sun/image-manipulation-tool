package main_test

import (
	"image-manipulation-tool"
	"math"
	"testing"
)

func TestGetTranslationMatrix(t *testing.T) {
	x := 3.0
	y := 5.0
	receivedMatrix := main.GetTranslationMatrix(x, y)
	expectedMatrix := [][]float64{
		{1, 0, x},
		{0, 1, y},
		{0, 0, 1},
	}
	for i := range receivedMatrix {
		for j := range receivedMatrix[i] {
			if receivedMatrix[i][j] != expectedMatrix[i][j] {
				t.Errorf(main.RError()+" expected translation receivedMatrix element at [%d][%d] to be %f, but got %f", i, j, expectedMatrix[i][j], receivedMatrix[i][j])
			}
		}
	}
}

func TestGetResizeMatrix(t *testing.T) {
	x := 3.0
	y := 5.0
	receivedMatrix := main.GetResizeMatrix(x, y)
	expectedMatrix := [][]float64{
		{1 / x, 0, 0},
		{0, 1 / y, 0},
		{0, 0, 1},
	}
	for i := range receivedMatrix {
		for j := range receivedMatrix[i] {
			if receivedMatrix[i][j] != expectedMatrix[i][j] {
				t.Errorf(main.RError()+" expected translation receivedMatrix element at [%d][%d] to be %f, but got %f", i, j, expectedMatrix[i][j], receivedMatrix[i][j])
			}
		}
	}
}

func TestGetMirrorHMatrix(t *testing.T) {
	receivedMatrix := main.GetMirrorHMatrix()
	expectedMatrix := [][]float64{
		{-1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	for i := range receivedMatrix {
		for j := range receivedMatrix[i] {
			if receivedMatrix[i][j] != expectedMatrix[i][j] {
				t.Errorf(main.RError()+" expected translation receivedMatrix element at [%d][%d] to be %f, but got %f", i, j, expectedMatrix[i][j], receivedMatrix[i][j])
			}
		}
	}
}

func TestGetMirrorVMatrix(t *testing.T) {
	receivedMatrix := main.GetMirrorVMatrix()
	expectedMatrix := [][]float64{
		{1, 0, 0},
		{0, -1, 0},
		{0, 0, 1},
	}
	for i := range receivedMatrix {
		for j := range receivedMatrix[i] {
			if receivedMatrix[i][j] != expectedMatrix[i][j] {
				t.Errorf(main.RError()+" expected translation receivedMatrix element at [%d][%d] to be %f, but got %f", i, j, expectedMatrix[i][j], receivedMatrix[i][j])
			}
		}
	}
}

func TestGetRotationMatrix(t *testing.T) {
	x := 3.0
	receivedMatrix := main.GetRotationMatrix(x)
	expectedMatrix := [][]float64{
		{math.Cos(x), math.Sin(x), 0},
		{-math.Sin(x), math.Cos(x), 0},
		{0, 0, 1},
	}
	for i := range receivedMatrix {
		for j := range receivedMatrix[i] {
			if receivedMatrix[i][j] != expectedMatrix[i][j] {
				t.Errorf(main.RError()+" expected translation receivedMatrix element at [%d][%d] to be %f, but got %f", i, j, expectedMatrix[i][j], receivedMatrix[i][j])
			}
		}
	}
}
