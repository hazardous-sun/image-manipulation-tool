package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type DisplayedImages struct {
	originPath    string
	originalImage image.Image
	previewPath   string
	previewImage  image.Image
}

func loadImage(path string) (image.Image, error) {
	// Open the file
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Check if the file is a supported image
	fileExt := filepath.Ext(path)
	switch fileExt {
	case ".jpg", ".jpeg":
		image.RegisterFormat("jpeg", "jepg", jpeg.Decode, jpeg.DecodeConfig)
		return jpeg.Decode(file)
	case ".png":
		image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
		return png.Decode(file)
	default:
		return image.Image(nil), fmt.Errorf("unsupported image format: '%s'", fileExt)
	}
}
