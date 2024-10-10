package fyne_project

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image-manipulation-tool/fyne_project/models"
)

func Build(a fyne.App) {
	project := models.NewProject()
	w := a.NewWindow("Image Manipulation Tool")
	w.Resize(fyne.NewSize(800, 600))

	w.SetMainMenu(
		fyne.NewMainMenu(
			fyne.NewMenu("File",
				fyne.NewMenuItem("New", func() {}),
				fyne.NewMenuItem("Close", func() {}),
			),
			fyne.NewMenu("Help",
				fyne.NewMenuItem("About", func() {}),
			),
		),
	)

	originalImage := project.GetOriginal()
	originalImageCanvas := canvas.NewImageFromImage(originalImage)
	originalImageCanvas.FillMode = canvas.ImageFillOriginal

	previewImage := project.GetPreview()
	previewImageCanvas := canvas.NewImageFromImage(previewImage)
	previewImageCanvas.FillMode = canvas.ImageFillOriginal
	imgsCtr := container.NewBorder(
		nil, nil,
		originalImageCanvas,
		previewImageCanvas,
	)
	w.SetContent(imgsCtr)
	w.ShowAndRun()
}
