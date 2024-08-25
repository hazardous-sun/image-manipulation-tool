package main

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"image"
	color2 "image/color"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
)

/*
Sends a message to the JavaScript listener informing that both the original and preview images should be updated.
Useful when the extension for the image changes and the file cannot simply be overwritten.
*/
func setOriginPrev(app *App, path string) {
	println(path)

	createImage(path, true)
	createImage(path, false)

	println("mandando para o JAVASCRIPT")
	notifyImagesChange(app, path, true)
}

func notifyImagesChange(app *App, path string, both bool) {
	if both {
		runtime.EventsEmit(app.ctx, "set-origin-prev", map[string]interface{}{
			"fileExt": filepath.Ext(path),
		})
	} else {
		runtime.EventsEmit(app.ctx, "set-prev", map[string]interface{}{
			"fileExt": filepath.Ext(path),
		})
	}
}

func removeAllFiles(dirPath string) error {
	// Get a list of all files in the directory
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("error reading directory: %w", err)
	}

	// Iterate through each file and remove it
	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())
		if err := os.Remove(filePath); err != nil {
			return fmt.Errorf("error removing file: %w", err)
		}
	}

	return nil
}

/*
Collects a file from the system and replicates it at "frontend/src/assets/temp", where it can be accessed by the
frontend.
*/
func createImage(originalPath string, origin bool) {
	var path string

	if origin {
		fileCount, err := countFiles("frontend/src/assets/temp/origin/")

		if err != nil {
			println(err.Error())
			return
		}

		path = "frontend/src/assets/temp/origin/" + string(rune(fileCount)) + filepath.Ext(originalPath)
	} else {
		fileCount, err := countFiles("frontend/src/assets/temp/origin/")

		if err != nil {
			println(err.Error())
			return
		}

		path = "frontend/src/assets/temp/prev/" + string(rune(fileCount)) + filepath.Ext(originalPath)
	}

	// Load original file
	originalFile, err := os.Open(originalPath)

	if err != nil {
		println("Error during file opening:", err.Error())
	}
	defer originalFile.Close()

	err = copyFile(path, originalFile)

	if err != nil {
		println("Error during image saving:", err.Error())
	}
}

/*
Returns the file count of a directory.
*/
func countFiles(path string) (int, error) {
	files, err := os.ReadDir(path)

	if err != nil {
		return -1, err
	}

	return len(files), nil
}

/*
Copies the content from a file to a new one.
*/
func copyFile(path string, content *os.File) error {
	// Create new file
	destFile, err := os.Create(path)

	if err != nil {
		return err
	}
	defer destFile.Close()

	// Copy content to the created file
	_, err = io.Copy(destFile, content)
	if err != nil {
		return err
	}

	return nil
}

/*
Decodes an image file and transform it into an "image.Image".
*/
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
		image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
		return jpeg.Decode(file)
	case ".png":
		image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
		return png.Decode(file)
	default:
		return image.Image(nil), fmt.Errorf("unsupported image format: '%s'", fileExt)
	}
}

/*
Encodes an image into a file.
*/
func saveImage(path string, fileExt string, img image.Image) error {
	//err := removeAllFiles("frontend/src/assets/temp")
	//
	//if err != nil {
	//	return err
	//}

	file, err := os.Create(path)

	if err != nil {
		return err
	}
	defer file.Close()

	switch fileExt {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(file, img, nil)
		if err != nil {
			return err
		}
	case ".png":
		err = png.Encode(file, img)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported image format: '%s'", fileExt)
	}
	return nil
}

// Filters -------------------------------------------------------------------------------------------------------------

// Grayscale
/*
Applies the grayscale filter to an image.
*/
func filterGrayScale(img image.Image) image.Image {
	grayImage := image.NewGray(img.Bounds())
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			grayIntensity := (r + g + b) / 3
			grayImage.Set(x, y, color2.Gray{
				Y: uint8(grayIntensity),
			})
		}
	}
	fmt.Println(grayImage.GrayAt(10, 10))
	return grayImage
}

// ---------------------------------------------------------------------------------------------------------------------
