package fyne_project

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"image-manipulation-tool/fyne_project/models"
)

func Build(a fyne.App) {
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

	imgsCtr := initializeImgsCtr()
	sideBar := initializeSideBar()

	appCtr := container.NewBorder(
		nil, nil,
		imgsCtr,
		sideBar,
	)

	w.SetContent(appCtr)
	w.ShowAndRun()
}

func initializeImgsCtr() fyne.CanvasObject {
	// initialize the Project instance
	project := models.NewProject()

	// initialize the original image canvas
	originalImage := project.GetOriginal()
	originalImageCanvas := canvas.NewImageFromImage(originalImage)
	originalImageCanvas.FillMode = canvas.ImageFillOriginal

	// build a container for the original image
	originalImageCtr := container.NewBorder(
		widget.NewLabel("Original image"),
		nil, nil, nil,
		originalImageCanvas,
	)

	// initialize the preview image canvas
	previewImage := project.GetPreview()
	previewImageCanvas := canvas.NewImageFromImage(previewImage)
	previewImageCanvas.FillMode = canvas.ImageFillOriginal

	// build a container for the preview image
	previewImageCtr := container.NewBorder(
		widget.NewLabel("Preview image"),
		nil, nil, nil,
		previewImageCanvas,
	)

	return container.NewBorder(
		nil, nil,
		originalImageCtr,
		previewImageCtr,
	)
}

func initializeSideBar() fyne.CanvasObject {
	btnsArr := []models.EditButton{
		models.NewEditButton("Resize", func() {}),
		models.NewEditButton("Rotate", func() {}),
		models.NewEditButton("Translate", func() {}),
		models.NewEditButton("Horizontal mirroring", func() {}),
		models.NewEditButton("Vertical mirroring", func() {}),
	}

	btnsData := binding.NewUntypedList()
	for _, btn := range btnsArr {
		_ = btnsData.Append(btn)
	}

	sideBar := widget.NewAccordion(
		widget.NewAccordionItem(
			"Geometric trasnformations",
			widget.NewListWithData(
				btnsData,
				func() fyne.CanvasObject {
					return widget.NewButton("", func() {})
				},
				func(di binding.DataItem, object fyne.CanvasObject) {
					objBtn := object.(*widget.Button)
					temp, _ := di.(binding.Untyped).Get()
					diBtn := temp.(models.EditButton)
					objBtn.SetText(diBtn.Text)
					objBtn.OnTapped = diBtn.Behavior
				},
			),
		),
		widget.NewAccordionItem("Filters", widget.NewLabel("Teste")),
		widget.NewAccordionItem("Mathematical morphology", widget.NewLabel("Teste")),
		widget.NewAccordionItem("Feature extraction", widget.NewLabel("Teste")),
	)

	return container.NewBorder(
		nil, nil, nil,
		sideBar,
	)
}
