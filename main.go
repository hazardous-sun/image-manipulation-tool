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

	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	FileMenu.AddText("&Open", keys.CmdOrCtrl("o"), func(_ *menu.CallbackData) {
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
	FileMenu.AddText("Exit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})

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
