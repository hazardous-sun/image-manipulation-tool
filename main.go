package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	goRuntime "runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Set menu
	AppMenu := setMenuBar(app)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Image Manipulation Tool",
		Width:  1024,
		Height: 768,
		Menu:   AppMenu,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 35, G: 35, B: 35, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func setMenuBar(app *App) *menu.Menu {
	AppMenu := menu.NewMenu()
	setFileMenu(app, AppMenu)
	setGeoTransformMenu(app, AppMenu)
	return AppMenu
}

func setFileMenu(app *App, AppMenu *menu.Menu) {
	// File
	FileMenu := AppMenu.AddSubmenu("File")
	// -- Open image
	FileMenu.AddText("Open", keys.CmdOrCtrl("o"), func(_ *menu.CallbackData) {
		cwd, err := os.Getwd()

		if err != nil {
			println("Error:", err.Error())
			if goRuntime.GOOS == "windows" {
				cwd = "%USERPROFILE%"
			} else {
				cwd = "~"
			}
		}

		imagePath, err := runtime.OpenFileDialog(
			app.ctx,
			runtime.OpenDialogOptions{
				DefaultDirectory:           cwd,
				DefaultFilename:            "",
				Title:                      "Select image",
				Filters:                    nil,
				ShowHiddenFiles:            false,
				CanCreateDirectories:       false,
				ResolvesAliases:            false,
				TreatPackagesAsDirectories: false,
			},
		)

		if err != nil {
			println("Error:", err.Error())
			return
		}

		println(imagePath)
	})
	FileMenu.AddSeparator()
	// -- Save image
	FileMenu.AddSeparator()

	// -- About
	FileMenu.AddText("About", keys.CmdOrCtrl("f1"), func(_ *menu.CallbackData) {
		runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
			Type:          runtime.InfoDialog,
			Title:         "About",
			Message:       "This is an image manipulation tool written in Go using Wails framework to build the frontend.",
			DefaultButton: "Back",
		})
	})
	// -- Exit
	FileMenu.AddSeparator()
	FileMenu.AddText("Exit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})
}

func setGeoTransformMenu(app *App, AppMenu *menu.Menu) {
	// Geometric Transformations
	GeoTransformMenu := AppMenu.AddSubmenu("Geometric Transformations")
	// -- Translate
	GeoTransformMenu.AddText("Translate", keys.Key("t"), func(_ *menu.CallbackData) {})
	GeoTransformMenu.AddSeparator()
	// -- Rotate
	GeoTransformMenu.AddText("Rotate", keys.Key("r"), func(_ *menu.CallbackData) {})
	GeoTransformMenu.AddSeparator()
	// -- Horizontal mirroring
	GeoTransformMenu.AddText("Horizontal Mirroring", keys.CmdOrCtrl("h"), func(_ *menu.CallbackData) {})
	GeoTransformMenu.AddSeparator()
	// -- Vertical mirroring
	GeoTransformMenu.AddText("Vertical Mirroring", keys.Key("v"), func(_ *menu.CallbackData) {})
	GeoTransformMenu.AddSeparator()
	// -- Resize
	GeoTransformMenu.AddText("Resize", keys.Key("w"), func(_ *menu.CallbackData) {})
}
