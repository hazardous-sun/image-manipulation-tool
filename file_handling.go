package main

import (
	"bytes"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
)

// File handling -------------------------------------------------------------------------------------------------------

// Removes all files in the desired path.
func removeAllFiles(dirPath string) error {
	// Get a list of all files in the directory
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf(pError()+" reading directory: %w", err)
	}

	// Iterate through each file and remove it
	for _, file := range files {
		filePath := filepath.Join(
			dirPath,
			file.Name(),
		)
		if err := os.Remove(filePath); err != nil {
			return fmt.Errorf(pError()+" removing file: %w", err)
		}
	}

	return nil
}

// Returns the file count of a directory.
func countFiles(path string) (int, error) {
	files, err := os.ReadDir(path)

	if err != nil {
		return -1, err
	}

	return len(files), nil
}

// Copies the content from a file to a new one.
func copyFile(path string, content *os.File) error {
	// Create new file
	destFile, err := os.Create(path)

	if err != nil {
		return err
	}

	defer func(destFile *os.File) {
		err := destFile.Close()
		if err != nil {
			println(pError()+" closing file '%s': %w", path, err)
		}
	}(destFile)

	// Copy content to the created file
	_, err = io.Copy(destFile, content)
	if err != nil {
		return err
	}

	return nil
}

// Image handling ------------------------------------------------------------------------------------------------------

// Informs the frontend that the images need to be reloaded.
func notifyImagesChange(app *App, path string, both bool) {
	// Check if both images will change
	if both {
		// If both will change, origin and prev directories will be cleared, so the file can be saved as "0.{file ext}"
		runtime.EventsEmit(app.ctx, "set-origin-prev", map[string]interface{}{
			"path": "0" + filepath.Ext(path),
		})
	} else {
		// If only the preview image will change, it is required to check how many prev images already exist in the dir
		fileCount, err := countFiles(filepath.Join(
			"frontend",
			"src",
			"assets",
			"temp",
			"prev",
		))

		if err != nil {
			println(err.Error())
			return
		}

		if fileCount > 1 {
			fileCount -= 1
		}

		runtime.EventsEmit(app.ctx, "set-prev", map[string]interface{}{
			"path": fmt.Sprintf("%d", fileCount) + filepath.Ext(path),
		})
	}
}

// Sends a message to the JavaScript listener informing that both the original and preview images should be updated.
// Useful when the extension for the image changes and the file cannot simply be overwritten.
func setOriginPrev(app *App, path string) {
	// Clean origin and prev directories
	err := removeAllFiles(filepath.Join(
		"frontend",
		"src",
		"assets",
		"temp",
		"origin",
	))

	if err != nil {
		println(pError()+" while trying to clean origin images directory: %w", err)
	}

	err = removeAllFiles(filepath.Join(
		"frontend",
		"src",
		"assets",
		"temp",
		"prev",
	))

	if err != nil {
		println(pError()+" while trying to clean preview images directory: %w", err)
	}

	// Create images
	createImage(path, true)
	createImage(path, false)

	// Passes the path in frontend/src/assets/temp/...
	notifyImagesChange(app, filepath.Ext(path), true)
}

// Collects a file from the system and replicates it at "frontend/src/assets/temp", where it can be accessed by the
// frontend.
func createImage(originalPath string, origin bool) {
	var newImagePath string

	if origin {
		newImagePath = filepath.Join(
			"frontend",
			"src",
			"assets",
			"temp",
			"origin",
			"0"+filepath.Ext(originalPath),
		)
	} else {
		fileCount, err := countFiles(filepath.Join(
			"frontend",
			"src",
			"assets",
			"temp",
			"prev",
		))

		if err != nil {
			println(err.Error())
			return
		}

		if fileCount > 0 {
			fileCount = fileCount - 1
		}

		newImagePath = filepath.Join(
			"frontend",
			"src",
			"assets",
			"temp",
			"prev",
			fmt.Sprintf("%d", fileCount)+filepath.Ext(originalPath))
	}

	// Load original file
	originalFile, err := os.Open(originalPath)

	if err != nil {
		println(pError()+" when file opening:", err.Error())
	}
	defer func(originalFile *os.File) {
		err := originalFile.Close()
		if err != nil {
			println(pError()+" when closing file '%s': %w", originalPath, err)
		}
	}(originalFile)

	err = copyFile(newImagePath, originalFile)

	if err != nil {
		println(pError()+" when image saving:", err.Error())
	}
}

// Decodes an image file and transform it into an "image.Image".
func loadImage(path string) (image.Image, error) {
	// Open the file
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			println(pError()+" when closing file '%s': %w", path, err)
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

func loadImageToBytes(path string) ([]byte, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf(pError()+" when opening file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			println(pError()+" when closing file '%s': %w", path, err)
		}
	}(file)

	img, _, err := image.Decode(file)

	if err != nil {
		return nil, fmt.Errorf(pError()+" when decoding file: %w", err)
	}

	buf := &bytes.Buffer{}
	err = png.Encode(buf, img)

	if err != nil {
		return nil, fmt.Errorf(pError()+" when encoding file to bytes: %w", err)
	}

	return buf.Bytes(), nil
}

// Encodes an image into a file.
func saveImage(path string, fileExt string, img image.Image) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			println(pError()+" when closing file '%s': %w", path, err)
		}
	}(file)

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

// Prints a red error message
func pError() string {
	return "\033[31merror:\033[0m"
}
