package main

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"path/filepath"
	goRuntime "runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

// Build : Holds the application instance, the menu and the options used for the display.
type Build struct {
	AppInstance        *App
	AppOptions         options.App
	TempDirInitialized bool
}

// Constructs the application instance, the menu and the options used for the display.
func (b Build) build() (Build, error) {
	appInstance, appMenu := setApp()
	appOptions := setOptions(appInstance, appMenu)
	err := initializeTemporaryDir()

	if err != nil {
		return Build{}, err
	}

	return Build{
		AppInstance:        appInstance,
		AppOptions:         appOptions,
		TempDirInitialized: true,
	}, nil
}

// ---------------------------------------------------------------------------------------------------------------------

// Run : Initializes the application.
func Run(build Build) {
	// Initialize the application with the chosen appOptions
	err := wails.Run(&build.AppOptions)

	if err != nil {
		println(pError()+":", err.Error())
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Passes the menu to the app instance.
func setApp() (*App, *menu.Menu) {
	// Create an instance of the app structure
	app := NewApp()

	// Set menu
	AppMenu := setMenu(app)

	return app, AppMenu
}

// ---------------------------------------------------------------------------------------------------------------------

// Returns the options used for building the application.
func setOptions(app *App, AppMenu *menu.Menu) options.App {
	icon, err := loadImageToBytes(filepath.Join(
		"build",
		"appicon.png",
	))

	if err != nil {
		println(err.Error())
		icon = nil
	}

	return options.App{
		Title:     "Image Manipulation Tool",
		Height:    1000,
		MinHeight: 500,
		Width:     1200,
		MinWidth:  500,
		Menu:      AppMenu,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		OnShutdown: func(ctx context.Context) {
			err := removeTemporaryDir()

			if err != nil {
				println(pError()+":", err.Error())
			}
		},
		Bind: []interface{}{
			app,
		},
		Linux: &linux.Options{
			Icon:                icon,
			WindowIsTranslucent: false,
			Messages:            nil,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyOnDemand,
			ProgramName:         "Image Manipulation Tool",
		},
	}
}

// Temp dir ------------------------------------------------------------------------------------------------------------

// Creates a directory under "frontend/src/assets/temp" that holds a copy of the original image, along with the preview
// images.
func initializeTemporaryDir() error {
	dir := filepath.Join(
		"frontend",
		"src",
		"assets",
		"temp",
	)
	err := os.Mkdir(dir, os.ModePerm)

	if err != nil {
		if os.IsExist(err) {
			return nil
		} else {
			println(pError()+" while creating the temp directory:", err.Error())
			return err
		}
	}

	err = os.Mkdir(filepath.Join(
		dir,
		"origin",
	), os.ModePerm)

	if err != nil {
		if os.IsExist(err) {
			return nil
		} else {
			println(pError()+" while creating the temp origin directory:", err.Error())
			return err
		}
	}

	err = os.Mkdir(filepath.Join(
		dir,
		"prev",
	), os.ModePerm)

	if err != nil {
		if os.IsExist(err) {
			return nil
		} else {
			println(pError()+" while creating the temp prev directory:", err.Error())
			return err
		}
	}

	return nil
}

// Removes the temporary directory that holds the copy of the original and preview images.
func removeTemporaryDir() error {
	dir := filepath.Join(
		"frontend",
		"src",
		"assets",
		"temp",
	)
	err := os.RemoveAll(dir)
	if err != nil {
		return err
	}
	return nil
}

// Menu ----------------------------------------------------------------------------------------------------------------

// Creates the items in the menu bar at the top of the application.
func setMenu(app *App) *menu.Menu {
	AppMenu := menu.NewMenu()
	setFileMenu(app, AppMenu)
	return AppMenu
}

// ---------- Menu items

// Sets the "File" menu at the top menu bar.
func setFileMenu(app *App, AppMenu *menu.Menu) {
	// File
	FileMenu := AppMenu.AddSubmenu("File")
	// -- Open image
	FileMenu.AddText("Open", keys.CmdOrCtrl("o"), func(_ *menu.CallbackData) {
		// Try to collect current working directory
		cwd, err := os.Getwd()

		if err != nil {
			println(pError()+" when CWD collection:", err.Error())
			if goRuntime.GOOS == "windows" {
				cwd = "%USERPROFILE%"
			} else {
				cwd = "~"
			}
		}

		path, err := runtime.OpenFileDialog(
			app.ctx,
			runtime.OpenDialogOptions{
				DefaultDirectory: cwd,
				DefaultFilename:  "",
				Title:            "Select image",
				Filters: []runtime.FileFilter{
					{
						DisplayName: "Image Files (*.gif, *.jpeg, *.jpg, *.png)",
						Pattern:     "*.gif;*.jpeg;*.jpg;*.png;",
					},
				},
				ShowHiddenFiles:            false,
				CanCreateDirectories:       false,
				ResolvesAliases:            false,
				TreatPackagesAsDirectories: false,
			},
		)

		if err != nil {
			println(pError()+" during dialog:", err.Error())
			return
		}

		setOriginPrev(app, path)
	})
	FileMenu.AddSeparator()
	// -- Save image
	FileMenu.AddText("Save", keys.CmdOrCtrl("s"), func(_ *menu.CallbackData) {
		cwd, err := os.Getwd()

		if err != nil {
			println(pError()+" during CWD collection:", err.Error())
			if goRuntime.GOOS == "windows" {
				cwd = "%USERPROFILE%"
			} else {
				cwd = "~"
			}
		}

		path, err := runtime.SaveFileDialog(
			app.ctx,
			runtime.SaveDialogOptions{
				DefaultDirectory:           cwd,
				DefaultFilename:            "",
				Title:                      "Save image",
				Filters:                    nil,
				ShowHiddenFiles:            true,
				CanCreateDirectories:       true,
				TreatPackagesAsDirectories: false,
			},
		)

		runtime.EventsEmit(app.ctx, "get-prev", nil)

		runtime.EventsOn(app.ctx, "receive-prev", func(optionalData ...interface{}) {
			println("path received from JS:", optionalData) // TODO fix this
			img, err := loadImage(path)

			if err != nil {
				println(pError()+" when loading image:", err.Error())
				return
			}

			err = saveImage(path, filepath.Ext(path), img)

			if err != nil {
				println(pError()+" when saving image:", err.Error())
				return
			}
		})
	})
	FileMenu.AddSeparator()
	// -- About
	FileMenu.AddText("About", keys.CmdOrCtrl("f1"), func(_ *menu.CallbackData) {
		_, err := runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
			Type:          runtime.InfoDialog,
			Title:         "About",
			Message:       "This is an image manipulation tool written in Go using Wails framework to Run the frontend.",
			DefaultButton: "Back",
		})

		if err != nil {
			println(pError()+" during dialog: %w", err)
		}
	})
	FileMenu.AddSeparator()
	// -- Exit
	FileMenu.AddText("Exit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})
}
