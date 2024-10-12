package image_editing

import (
	"image"
	"image/color"
	"math/rand"
	"testing"
)

func TestFilterGrayScale(t *testing.T) {
	originalImg := image.NewRGBA(image.Rect(0, 0, 256, 256))

	for x := 0; x < originalImg.Bounds().Dx(); x++ {
		for y := 0; y < originalImg.Bounds().Dy(); y++ {
			r := rand.Intn(256)
			g := rand.Intn(256)
			b := rand.Intn(256)
			originalImg.Set(x, y, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255})
		}
	}

	received := FilterGrayScale(originalImg)
	if received.ColorModel() != color.GrayModel {
		t.Errorf("expected color model %v, got %v", color.GrayModel, received.ColorModel())
	}
}
