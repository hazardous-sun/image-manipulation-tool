package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image"
	"image-manipulation-tool/file_handling"
	"image-manipulation-tool/image_editing"
	"image-manipulation-tool/models"
	"io"
	"strconv"
)

var originalImageCanvas *canvas.Image
var previewImageCanvas *canvas.Image

func Build(a fyne.App) {
	// initialize the Project instance
	project := models.NewProject()

	// initialize the main window
	w := a.NewWindow("Image Manipulation Tool")
	w.Resize(fyne.NewSize(800, 600))

	// initialize GUI's elements
	imgsCtr := initializeImgsCtr(project)
	sideBar := initializeSideBar(a, project)

	// pass the elements to a container
	appCtr := container.NewBorder(
		nil, nil, nil,
		sideBar,
		imgsCtr,
	)

	// -- initialize the main menu for the window
	appMenu := initializeAppMenu(w, project)
	w.SetMainMenu(appMenu)

	// set the container as the main window's content
	w.SetContent(appCtr)

	// raise the main window and run the application
	w.ShowAndRun()
}

func initializeImgsCtr(project *models.Project) fyne.CanvasObject {
	// initialize the original image canvas
	originalImage := project.GetOriginal()
	originalImageCanvas = canvas.NewImageFromImage(originalImage)
	originalImageCanvas.FillMode = canvas.ImageFillContain

	// initialize original image label
	originalImageLbl := widget.NewLabel("Original image")
	originalImageLbl.Alignment = fyne.TextAlignCenter
	originalImageLbl.TextStyle = fyne.TextStyle{Bold: true}

	// build a container for the original image
	originalImageCtr := container.NewBorder(
		originalImageLbl,
		nil, nil, nil,
		originalImageCanvas,
	)

	// initialize the preview image canvas
	previewImage := project.GetPreview()
	previewImageCanvas = canvas.NewImageFromImage(previewImage)
	previewImageCanvas.FillMode = canvas.ImageFillContain

	// initialize original image label
	previewImageLbl := widget.NewLabel("Original image")
	previewImageLbl.Alignment = fyne.TextAlignCenter
	previewImageLbl.TextStyle = fyne.TextStyle{Bold: true}

	// build a container for the preview image
	previewImageCtr := container.NewBorder(
		previewImageLbl,
		nil, nil, nil,
		previewImageCanvas,
	)

	return container.NewGridWithColumns(
		2,
		originalImageCtr,
		previewImageCtr,
	)
}

func initializeSideBar(a fyne.App, project *models.Project) fyne.CanvasObject {
	// geometric transformations
	geoTransfBtns := getBtns(
		[]*widget.Button{
			widget.NewButton("Resize", func() {
				// initialize new window
				w := a.NewWindow("Input values")
				w.Resize(fyne.NewSize(200, 100))
				w.SetFixedSize(true)

				// initialize X axis input
				xEntry := widget.NewEntry()
				xEntry.PlaceHolder = "1"
				xCtr := container.NewBorder(
					nil, nil,
					widget.NewLabel("X:"),
					nil,
					xEntry,
				)

				// initialize Y axis input
				yEntry := widget.NewEntry()
				yEntry.PlaceHolder = "1"
				yCtr := container.NewBorder(
					nil, nil,
					widget.NewLabel("Y:"),
					nil,
					yEntry,
				)

				// initialize the confirmation button
				confirmBtn := widget.NewButton(
					"Confirm",
					func() {
						// transform the inputted string in X into a float64
						x, err := strconv.ParseFloat(xEntry.Text, 64)

						if err != nil {
							dialog.ShowError(err, w)
							return
						}

						// transform the inputted string in Y into a float64
						y, err := strconv.ParseFloat(yEntry.Text, 64)

						if err != nil {
							dialog.ShowError(err, w)
							return
						}

						// collect the matrix
						matrix := image_editing.GetResizeMatrix(x, y)

						// run the transformation process
						img := image_editing.TransformImage(previewImageCanvas.Image, matrix)

						// inform the system to update the preview image
						updatePrevImage(img, project)
						w.Close()
					},
				)
				ctr := container.NewGridWithRows(3,
					xCtr,
					yCtr,
					confirmBtn,
				)
				w.SetContent(ctr)
				w.Show()
			}),
			widget.NewButton("Rotate", func() {
				fmt.Println("rotate")
			}),
			widget.NewButton("Translate", func() {}),
			widget.NewButton("Horizontal mirroring", func() {}),
			widget.NewButton("Vertical mirroring", func() {}),
		},
	)
	geoTransfList := getBtnsList(geoTransfBtns)

	// filters
	filtersBtns := getBtns(
		[]*widget.Button{
			widget.NewButton("Grayscale", func() {
				img := image_editing.FilterGrayScale(previewImageCanvas.Image)
				updatePrevImage(img, project)
			}),
			widget.NewButton("High fade", func() {}),
			widget.NewButton("Low fade", func() {}),
			widget.NewButton("Threshold", func() {}),
		},
	)
	filterList := getBtnsList(filtersBtns)

	// mathematical morphology
	mathMorphoBtns := getBtns(
		[]*widget.Button{
			widget.NewButton("Dilatation", func() {}),
			widget.NewButton("Erosion", func() {}),
			widget.NewButton("Opening", func() {}),
			widget.NewButton("Closing", func() {}),
		},
	)
	mathMorphoList := getBtnsList(mathMorphoBtns)

	// pass the buttons list to the accordion
	sideBar := widget.NewAccordion(
		widget.NewAccordionItem(
			"Geometric trasnformations",
			geoTransfList,
		),
		widget.NewAccordionItem(
			"Filters",
			filterList,
		),
		widget.NewAccordionItem(
			"Mathematical morphology",
			mathMorphoList,
		),
		widget.NewAccordionItem(
			"Feature extraction",
			widget.NewLabel("Nothing here yet"),
		),
	)

	return container.NewBorder(
		nil, nil, nil,
		sideBar,
	)
}

func getBtns(btns []*widget.Button) binding.UntypedList {
	btnsList := binding.NewUntypedList()
	for _, btn := range btns {
		_ = btnsList.Append(btn)
	}

	return btnsList
}

func getBtnsList(btns binding.UntypedList) *widget.List {
	return widget.NewListWithData(
		btns,
		func() fyne.CanvasObject {
			return widget.NewButton("", func() {})
		},
		func(di binding.DataItem, object fyne.CanvasObject) {
			objBtn := object.(*widget.Button)
			temp, _ := di.(binding.Untyped).Get()
			diBtn := temp.(*widget.Button)
			objBtn.SetText(diBtn.Text)
			objBtn.OnTapped = diBtn.OnTapped
		},
	)
}

func initializeAppMenu(w fyne.Window, project *models.Project) *fyne.MainMenu {
	return fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Open", func() {
				dialog.ShowFileOpen(
					func(r fyne.URIReadCloser, err error) {
						if err != nil {
							dialog.ShowError(err, w)
							return
						}

						if r == nil {
							dialog.ShowError(fmt.Errorf("no file selected"), w)
							return
						} else {
							content, err := io.ReadAll(r)

							if err != nil {
								dialog.ShowError(err, w)
								return
							}

							img, err := file_handling.LoadImageFromBytes(content)

							if err != nil {
								dialog.ShowError(err, w)
							}

							updateAllImages(img, project)
						}
					},
					w,
				)
			}),
			fyne.NewMenuItem("Save", func() {}),
		),
		fyne.NewMenu("Help",
			fyne.NewMenuItem("About", func() {
				dialog.ShowInformation(
					"About",
					"This is an image manipulation tool written in Go that uses the Fyne framework for "+
						"building the frontend.",
					w,
				)
			}),
			fyne.NewMenuItem("Preferences", func() {}),
		),
	)
}

func updateAllImages(img image.Image, project *models.Project) {
	project.LoadNewImage(img)
	originalImageCanvas.Image = img
	previewImageCanvas.Image = img
	refreshCanvas()
}

func updatePrevImage(img image.Image, project *models.Project) {
	project.AddPreviewImage(img)
	previewImageCanvas.Image = project.GetPreview()
	refreshCanvas()
}

func refreshCanvas() {
	originalImageCanvas.Refresh()
	previewImageCanvas.Refresh()
}
