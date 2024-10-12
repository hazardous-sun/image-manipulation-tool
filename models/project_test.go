package models

import (
	"fmt"
	"image"
	"testing"
)

func TestNewProject(t *testing.T) {
	initialImg := image.NewRGBA(image.Rect(0, 0, 512, 512))

	expectedVersions := 1
	expectedOriginalImage := initialImg
	expectedPreviewImage := initialImg

	received := NewProject()

	if received.versions != expectedVersions {
		panic(fmt.Sprintf("expected versions %d, got %d", expectedVersions, received.versions))
	}

	if received.originalImage.Bounds() != expectedOriginalImage.Bounds() {
		panic(fmt.Sprintf("expected original image bounds %v, got %v", expectedOriginalImage.Bounds(), received.originalImage.Bounds()))
	}

	if received.previewImage.Bounds() != expectedPreviewImage.Bounds() {
		panic(fmt.Sprintf("expected preview image bounds %v, got %v", expectedPreviewImage.Bounds(), received.previewImage.Bounds()))
	}
}

func TestProject_GetOriginal(t *testing.T) {
	expected := image.NewRGBA(image.Rect(0, 0, 512, 512))
	received := NewProject()

	if received.GetOriginal().Bounds() != expected.Bounds() {
		panic(fmt.Sprintf("expected original image bounds %v, got %v", expected, received.GetOriginal()))
	}

	if received.GetOriginal().ColorModel() != expected.ColorModel() {
		panic(fmt.Sprintf("expected original color model %v, got %v", expected.ColorModel(), received.GetOriginal().ColorModel()))
	}
}
