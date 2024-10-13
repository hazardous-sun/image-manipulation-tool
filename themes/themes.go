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
	"os/user"
)

// ThemeSelectionWindow :
// Raises a window where the user can toggle between themes.
func ThemeSelectionWindow(a fyne.App, settings *models.ThemeSettings) {
	var selectedTheme *models.CustomTheme

	// initialize window
	w := a.NewWindow("Choose Theme")
	w.Resize(fyne.NewSize(700, 500))

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
			theme := temp.(models.CustomTheme)
			lbl := object.(*widget.Label)
			lbl.SetText(theme.Name())
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
				a.Settings().SetTheme(selectedTheme)
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
			// raise a new file selection dialog
			dialog.ShowFileOpen(
				func(closer fyne.URIReadCloser, err error) {
					// get the path for the config file
					path := closer.URI().Path()
					// parse it into a new CustomTheme
					customTheme, err := models.NewCustomTheme(path)

					if err != nil {
						dialog.ShowError(err, w)
						return
					}

					// append the new CustomTheme to themes
					themes = append(themes, *customTheme)

					// append the new CustomTheme to themesDI
					_ = themesDI.Append(*customTheme)
				},
				w,
			)
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

// LoadThemes :
// Tries to load themes from "$HOME/image-manipulation-tool-themes" into the memory.
func LoadThemes(settings *models.ThemeSettings) {
	homeDir, err := user.Current()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(homeDir.HomeDir)
}
