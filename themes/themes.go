package themes

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image-manipulation-tool/models"
)

var customThemes []models.CustomTheme

func ThemeSelectionWindow(a fyne.App, settings *models.ThemeSettings) {
	var selectedTheme *models.TestTheme

	// initialize window
	w := a.NewWindow("Choose Theme")
	w.Resize(fyne.NewSize(300, 400))

	themes := settings.Themes()
	themesDI := binding.NewUntypedList()
	for _, v := range themes {
		_ = themesDI.Append(v)
	}

	themesList := widget.NewListWithData(
		themesDI,
		func() fyne.CanvasObject {
			lbl := widget.NewLabel("")
			return lbl
		},
		func(di binding.DataItem, object fyne.CanvasObject) {
			temp, _ := di.(binding.Untyped).Get()
			theme := temp.(models.TestTheme)
			lbl := object.(*widget.Label)
			lbl.SetText(theme.Name)
		},
	)
	themesList.OnSelected = func(id widget.ListItemID) {
		selectedTheme = &themes[id]
	}

	btn := widget.NewButton(
		"Confirm",
		func() {
			if selectedTheme != nil {
				a.Settings().SetTheme(selectedTheme.Theme)
			} else {
				dialog.ShowError(
					fmt.Errorf("no theme selected"),
					w,
				)
			}
		},
	)

	selectionCtr := container.NewBorder(
		nil,
		btn,
		nil, nil,
		themesList,
	)

	w.SetContent(selectionCtr)
	w.Show()
}
