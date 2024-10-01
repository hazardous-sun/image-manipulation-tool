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

func (a *App) Transform(path string, code string, x float64, y float64) {
	var matrix [][]float64
	switch code {
	case "translate":
		matrix = GetTranslationMatrix(x, y)
	case "resize":
		matrix = GetResizeMatrix(x, y)
	case "mirrorH":
		matrix = GetMirrorHMatrix()
	case "mirrorV":
		matrix = GetMirrorVMatrix()
	case "rotate":
		matrix = GetRotationMatrix(x)
	default:
		return
	}

	path = "frontend" + path[29:]
	img, err := loadImage(path)

	if err != nil {
		println(RError()+" while loading image:", err.Error())
		return
	}

	img = transformImage(img, matrix)
	fileExt := filepath.Ext(path)
	fileCount, err := countFiles("frontend/src/assets/temp/prev/")

	if err != nil {
		println(RError()+" counting files:", err.Error())
		return
	}

	path = "frontend/src/assets/temp/prev/" + fmt.Sprintf("%d", fileCount) + fileExt
	err = saveImage(path, fileExt, img)

	if err != nil {
		println(RError()+" saving image:", err.Error())
		return
	}

	notifyImagesChange(a, path, false)

	UnsavedProgress = true
}

// Filters -------------------------------------------------------------------------------------------------------------

// GrayScale : Applies the grayscale filter to the preview image.
func (a *App) GrayScale(path string) {
	// Collect the path to the preview image
	// This solution needs to be refactored to avoid bugs in a situation where wails will initialize using a port with
	// a length != 5 (the standard port wails uses is 34115)
	path = "frontend" + path[29:]
	img, err := loadImage(path)

	if err != nil {
		println(RError()+" loading image:", err.Error())
		return
	}

	img = filterGrayScale(img)
	fileExt := filepath.Ext(path)
	fileCount, err := countFiles("frontend/src/assets/temp/prev/")

	if err != nil {
		println(RError()+" counting files:", err.Error())
		return
	}

	path = "frontend/src/assets/temp/prev/" + fmt.Sprintf("%d", fileCount) + fileExt
	err = saveImage(path, fileExt, img)

	if err != nil {
		println(RError()+" saving image:", err.Error())
		return
	}

	notifyImagesChange(a, path, false)

	UnsavedProgress = true
}
