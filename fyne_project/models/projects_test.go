package models

import (
	"image"
	"testing"
)

func TestNewProject(t *testing.T) {
	expectedVersions := 1
	expectedCurrentVersion := 0
	expectedOriginalImageBounds := image.NewRGBA(image.Rect(0, 0, 512, 512)).Bounds()
	expectedPreviewImageLen := 1

	received := NewProject()

	if received.versions != expectedVersions {
		t.Errorf("Expected versions: %d, got: %d\n", expectedVersions, received.versions)
	}

	if received.currentVersion != expectedCurrentVersion {
		t.Errorf("Expected currentVersion: %d, got: %d\n", expectedCurrentVersion, received.currentVersion)
	}

	if received.originalImage.Bounds() != expectedOriginalImageBounds {
		t.Errorf("Expected originalImageBounds: %v, got: %v\n", expectedOriginalImageBounds, received.originalImage.Bounds())
	}

	if len(received.previewImage) != expectedPreviewImageLen {
		t.Errorf("Expected previewImageLen: %d, got: %d\n", expectedPreviewImageLen, len(received.previewImage))
	}
}
