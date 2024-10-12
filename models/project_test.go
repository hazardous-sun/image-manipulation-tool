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

func TestProject_GetPreview(t *testing.T) {
	expected := image.NewRGBA(image.Rect(0, 0, 512, 512))
	received := NewProject()

	if received.GetPreview().Bounds() != expected.Bounds() {
		panic(fmt.Sprintf("expected preview image bounds %v, got %v", expected.Bounds(), received.GetPreview()))
	}

	if received.GetPreview().ColorModel() != expected.ColorModel() {
		panic(fmt.Sprintf("expected preview color model %v, got %v", expected.ColorModel(), received.GetOriginal().ColorModel()))
	}
}

func TestProject_LoadNewImage(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 256, 256))

	expected := Project{
		versions:       1,
		originalImage:  img,
		previewImage:   img,
		previousStates: ChangesStack{},
		nextStates:     ChangesStack{},
	}

	received := NewProject()
	received.LoadNewImage(img)

	if received.versions != expected.versions {
		panic(fmt.Sprintf("expected versions %d, got %d", expected.versions, received.versions))
	}

	if received.GetOriginal().Bounds() != expected.GetOriginal().Bounds() {
		panic(fmt.Sprintf("expected original image bounds %v, got %v", expected.GetOriginal(), received.GetOriginal()))
	}

	if received.GetOriginal().ColorModel() != expected.GetOriginal().ColorModel() {
		panic(fmt.Sprintf("expected original color model %v, got %v", expected.GetOriginal().ColorModel(), received.GetOriginal().ColorModel()))
	}

	if received.GetPreview().Bounds() != expected.GetPreview().Bounds() {
		panic(fmt.Sprintf("expected preview image bounds %v, got %v", expected.GetPreview().Bounds(), received.GetPreview().Bounds()))
	}

	if received.GetPreview().ColorModel() != expected.GetPreview().ColorModel() {
		panic(fmt.Sprintf("expected preview color model %v, got %v", expected.GetPreview().ColorModel(), received.GetPreview().ColorModel()))
	}
}

func TestProject_AddPreviewImage(t *testing.T) {
	expectedLen := 2
	received := NewProject()
	received.AddPreviewImage(image.NewRGBA(image.Rect(0, 0, 256, 256)))

	if received.previousStates.Length() != expectedLen {
		panic(fmt.Sprintf("expected length %d, got %d", expectedLen, received.previousStates.Length()))
	}
}

func TestProject_PreviousPreviewImage(t *testing.T) {
	received := NewProject()
	expectedImage := received.GetOriginal()
	received.AddPreviewImage(image.NewRGBA(image.Rect(0, 0, 256, 256)))

	receivedPrevious, err := received.PreviousPreviewImage()

	if err != nil {
		t.Fatal(err)
	}

	if receivedPrevious.Bounds() != expectedImage.Bounds() {
		panic(fmt.Sprintf("expected previous image bounds %v, got %v", expectedImage.Bounds(), receivedPrevious.Bounds()))
	}
}
