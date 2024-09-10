package main

import (
	"context"
	"fmt"
	"path/filepath"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Geometric transformations -------------------------------------------------------------------------------------------

func (a *App) Transform(path string, code int, x float64, y float64) {
	var matrix [][]float64
	switch code {
	case 0:
		matrix = getTranslationMatrix(x, y)
	case 1:
		matrix = getResizeMatrix(x, y)
	case 2:
		matrix = getMirrorHMatrix()
	case 3:
		matrix = getMirrorVMatrix()
	case 4:
		matrix = getRotationMatrix(x)
	}

	println(matrix)

	path = "frontend" + path[29:]
	img, err := loadImage(path)

	if err != nil {
		println("Error while loading image:", err.Error())
	}

	img = transformImage(img, matrix)
	fileExt := filepath.Ext(path)
	fileCount, err := countFiles("frontend/src/assets/temp/prev/")

	if err != nil {
		println("Error counting files:", err.Error())
		return
	}

	path = "frontend/src/assets/temp/prev/" + fmt.Sprintf("%d", fileCount) + fileExt
	err = saveImage(path, fileExt, img)

	if err != nil {
		println("Error saving image:", err.Error())
		return
	}

	notifyImagesChange(a, path, false)
}

// Filters -------------------------------------------------------------------------------------------------------------

// GrayScale
/*
Applies the grayscale filter to the preview image.
*/
func (a *App) GrayScale(path string) {
	// Collect the path to the preview image
	// This solution needs to be refactored to avoid bugs in a situation where wails will initialize using a port with
	// a length != 5 (the standard port wails uses is 34115)
	path = "frontend" + path[29:]
	img, err := loadImage(path)

	if err != nil {
		println("Error loading image:", err.Error())
		return
	}

	img = filterGrayScale(img)
	fileExt := filepath.Ext(path)
	fileCount, err := countFiles("frontend/src/assets/temp/prev/")

	if err != nil {
		println("Error counting files:", err.Error())
		return
	}

	path = "frontend/src/assets/temp/prev/" + fmt.Sprintf("%d", fileCount) + fileExt
	err = saveImage(path, fileExt, img)

	if err != nil {
		println("Error saving image:", err.Error())
		return
	}

	notifyImagesChange(a, path, false)
}
