package file_handling

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// LoadImage :
// Decodes an image file and transform it into an "image.Image".
func LoadImage(path string) (image.Image, error) {
	// Open the file
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			println("error: when closing file '%s': %w", path, err)
		}
	}(file)

	// Check if the file is a supported image
	fileExt := filepath.Ext(path)
	switch fileExt {
	case ".jpg", ".jpeg":
		image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
		return jpeg.Decode(file)
	case ".png":
		image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
		return png.Decode(file)
	default:
		return image.Image(nil), fmt.Errorf("unsupported image format: '%s'", fileExt)
	}
}

// LoadImageToBytes :
// Decodes an image file and transform it into an array of bytes.
func LoadImageToBytes(path string) ([]byte, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("error: when opening file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			println("error: when closing file '%s': %w", path, err)
		}
	}(file)

	img, _, err := image.Decode(file)

	if err != nil {
		return nil, fmt.Errorf("error: when decoding file: %w", err)
	}

	buf := &bytes.Buffer{}
	err = png.Encode(buf, img)

	if err != nil {
		return nil, fmt.Errorf("error: when encoding file to bytes: %w", err)
	}

	return buf.Bytes(), nil
}

func LoadImageFromBytes(arr []byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(arr))

	if err != nil {
		return nil, fmt.Errorf("error: when decoding file to bytes: %w", err)
	}

	return img, nil
}
