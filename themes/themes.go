package themes

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image-manipulation-tool/models"
)

// ThemeSelectionWindow :
// Raises a window where the user can toggle between themes.
func ThemeSelectionWindow(a fyne.App, settings *models.ThemeSettings) {
	var selectedTheme *models.TestTheme

	// initialize window
	w := a.NewWindow("Choose Theme")
	w.Resize(fyne.NewSize(300, 400))

	// collect themes from ThemeSettings
	themes := settings.Themes()
	themesDI := binding.NewUntypedList()
	for _, v := range themes {
		_ = themesDI.Append(v)
	}

	// initialize the list that displays the themes
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

	// initialize the buttons
	confirmBtn := widget.NewButton(
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
	addThemeBtn := widget.NewButton(
		"Add theme",
		func() {
			fmt.Println("add new theme")
		},
	)

	// initialize the container that holds the buttons
	btnsCtr := container.NewGridWithRows(
		4,
		layout.NewSpacer(),
		confirmBtn,
		addThemeBtn,
		layout.NewSpacer(),
	)

	// initialize the container that will be the window's content
	selectionCtr := container.NewBorder(
		nil, nil, nil,
		btnsCtr,
		themesList,
	)

	w.SetContent(selectionCtr)
	w.Show()
}
