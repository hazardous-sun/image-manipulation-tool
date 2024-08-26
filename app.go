package main

import (
	"context"
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

// GrayScale
/*
Applies the grayscale filter to the preview image.
*/
func (a *App) GrayScale(path string) {
	// TODO FIX THE PATH PASSED TO THE FRONTEND
	println("\n\n\nENTERED GrayScale()")
	path = "frontend" + path[29:]
	img, err := loadImage(path)

	if err != nil {
		println("Error loading image:", err.Error())
		return
	}

	img = filterGrayScale(img)
	fileExt := filepath.Ext(path)
	path = "frontend/src/assets/temp/prev" + fileExt
	err = saveImage(path, fileExt, img)

	if err != nil {
		println("Error saving image:", err.Error())
		return
	}

	notifyImagesChange(a, path, false)
}
