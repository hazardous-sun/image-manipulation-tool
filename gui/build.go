package gui

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
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

var originalImageCanvas *canvas.Image
var previewImageCanvas *canvas.Image
var currentVersion int
var lblCount *widget.Label

func Build(a fyne.App) {
	// initialize the Project instance
	project := models.NewProject()

	// initialize the main window
	w := a.NewWindow("Image Manipulation Tool")
	w.Resize(fyne.NewSize(800, 600))

	// initialize GUI's elements
	imgsCtr := initializeImgsCtr(project)
	sideBar := initializeSideBar(a, project)
	versionsCtr := initializeVersionsCtr(w, project)

	// creating a container for the images and versionsCtr
	displayCtr := container.NewBorder(
		nil,
		versionsCtr,
		nil, nil,
		imgsCtr,
	)

	// pass the elements to a container
	appCtr := container.NewBorder(
		nil, nil, nil,
		sideBar,
		displayCtr,
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

	// geometric transformations
	geoTransfBtns := getBtns(
		[]*widget.Button{
			widget.NewButton("Resize", func() {
				// initialize new window
				w := a.NewWindow("Input values")
				w.Resize(fyne.NewSize(200, 100))
				w.SetFixedSize(true)

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
						project.AddPreviewImage(img)
						updateAllImages(img, project)
						updateLblCount(1)
						w.Close()
					},
				)

				// pass the values to the container
				ctr := container.NewGridWithRows(3,
					xCtr,
					yCtr,
					confirmBtn,
				)
				w.SetContent(ctr)
				w.Show()
			}),
			widget.NewButton("Rotate", func() {
				// initialize new window
				w := a.NewWindow("Input values")
				w.Resize(fyne.NewSize(200, 100))
				w.SetFixedSize(true)

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

						// collect the matrix
						matrix := image_editing.GetRotationMatrix(x)

						// run the transformation process
						img := image_editing.TransformImage(previewImageCanvas.Image, matrix)

						// inform the system to update the preview image
						project.AddPreviewImage(img)
						updateAllImages(img, project)
						updateLblCount(1)
						w.Close()
					},
				)

				// pass the values to the container
				ctr := container.NewGridWithRows(2,
					xCtr,
					confirmBtn,
				)
				w.SetContent(ctr)
				w.Show()
			}),
			widget.NewButton("Translate", func() {
				// initialize new window
				w := a.NewWindow("Input values")
				w.Resize(fyne.NewSize(200, 100))
				w.SetFixedSize(true)

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
						matrix := image_editing.GetTranslationMatrix(x, y)

						// run the transformation process
						img := image_editing.TransformImage(previewImageCanvas.Image, matrix)

						// inform the system to update the preview image
						project.AddPreviewImage(img)
						updateAllImages(img, project)
						updateLblCount(1)
						w.Close()
					},
				)

				// pass the values to the container
				ctr := container.NewGridWithRows(3,
					xCtr,
					yCtr,
					confirmBtn,
				)
				w.SetContent(ctr)
				w.Show()
			}),
			widget.NewButton("Horizontal mirroring", func() {
				w := a.NewWindow("Input values")
				w.Resize(fyne.NewSize(200, 100))
				w.SetFixedSize(true)
				matrix := image_editing.GetMirrorHMatrix()
				img := image_editing.TransformImage(previewImageCanvas.Image, matrix)
				project.AddPreviewImage(img)
				updateAllImages(img, project)
				updateLblCount(1)
				w.Close()
			}),
			widget.NewButton("Vertical mirroring", func() {
				w := a.NewWindow("Input values")
				w.Resize(fyne.NewSize(200, 100))
				w.SetFixedSize(true)
				matrix := image_editing.GetMirrorVMatrix()
				img := image_editing.TransformImage(previewImageCanvas.Image, matrix)
				project.AddPreviewImage(img)
				updateAllImages(img, project)
				updateLblCount(1)
				w.Close()
			}),
		},
	)
	geoTransfList := getBtnsList(geoTransfBtns)

	// filters
	filtersBtns := getBtns(
		[]*widget.Button{
			widget.NewButton("Grayscale", func() {
				img := image_editing.FilterGrayScale(previewImageCanvas.Image)
				project.AddPreviewImage(img)
				updateAllImages(img, project)
				updateLblCount(1)
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

func initializeVersionsCtr(w fyne.Window, project *models.Project) fyne.CanvasObject {
	// label
	versionsCount := strconv.Itoa(currentVersion)
	lblCount = widget.NewLabel(fmt.Sprintf("Version: " + versionsCount))
	lblCount.Alignment = fyne.TextAlignCenter

	// buttons
	undoBtn := widget.NewButton(
		"Previous",
		func() {
			img, err := project.PreviousPreviewImage()

			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			updateLblCount(-1)
			updateAllImages(img, project)
		},
	)
	redoBtn := widget.NewButton(
		"Next",
		func() {
			img, err := project.NextPreviewImage()

			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			updateAllImages(img, project)
			updateLblCount(1)
		},
	)

	// container
	toolBar := container.NewGridWithColumns(
		3,
		undoBtn,
		lblCount,
		redoBtn,
	)
	return toolBar
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

							updateAllImagesNewProject(img, project)
						}
					},
					w,
				)
			}),
			fyne.NewMenuItem("Save", func() {
				dialog.ShowFileSave(
					func(closer fyne.URIWriteCloser, err error) {
						path := closer.URI().Path()
						file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModePerm)

						if err != nil {
							dialog.ShowError(err, w)
							return
						}

						previewImage := project.GetPreview()
						switch filepath.Ext(path) {
						case ".png":
							err = png.Encode(file, previewImage)
							if err != nil {
								dialog.ShowError(err, w)
							}
						case ".jpg", ".jpeg":
							err = jpeg.Encode(file, previewImage, &jpeg.Options{Quality: 100})
							if err != nil {
								dialog.ShowError(err, w)
							}
						default:
							dialog.ShowError(fmt.Errorf("unsupported file type"), w)
						}
					},
					w,
				)
			}),
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

func updateAllImagesNewProject(img image.Image, project *models.Project) {
	updateLblCount(-currentVersion)
	project.LoadNewImage(img)
	originalImageCanvas.Image = project.GetOriginal()
	previewImageCanvas.Image = project.GetPreview()
	refreshCanvas()
}

func updateAllImages(img image.Image, project *models.Project) {
	originalImageCanvas.Image = project.GetOriginal()
	previewImageCanvas.Image = project.GetPreview()
	refreshCanvas()
}

func updatePrevImage(img image.Image, project *models.Project) {
	previewImageCanvas.Image = project.GetPreview()
	refreshCanvas()
}

func refreshCanvas() {
	originalImageCanvas.Refresh()
	previewImageCanvas.Refresh()
}

func updateLblCount(val int) {
	currentVersion += val
	versionsCount := strconv.Itoa(currentVersion)
	lblCount.SetText(fmt.Sprintf("Version: " + versionsCount))
}
