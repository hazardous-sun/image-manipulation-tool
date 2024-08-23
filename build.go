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
	"io"
	"os"
	"path/filepath"
	goRuntime "runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

type Build struct {
	AppInstance        *App
	AppOptions         options.App
	TempDirInitialized bool
}

func (b Build) build() (Build, error) {
	appInstance, appMenu := setApp()
	appOptions := setOptions(appInstance, appMenu)
	err := initializeTemporaryDir()

	if err != nil {
		return Build{}, err
	}

	appImages = DisplayedImages{}

	return Build{
		AppInstance:        appInstance,
		AppOptions:         appOptions,
		TempDirInitialized: true,
	}, nil
}

func Run(build Build) {
	// Initialize the application with the chosen appOptions
	err := wails.Run(&build.AppOptions)

	if err != nil {
		println("Error:", err.Error())
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func setApp() (*App, *menu.Menu) {
	// Create an instance of the app structure
	app := NewApp()

	// Set menu
	AppMenu := setMenu(app)

	return app, AppMenu
}

// ---------------------------------------------------------------------------------------------------------------------

func setOptions(app *App, AppMenu *menu.Menu) options.App {
	return options.App{
		Title:     "Image Manipulation Tool",
		Height:    1000,
		MinHeight: 1000,
		Width:     1200,
		MinWidth:  800,
		Menu:      AppMenu,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		OnShutdown: func(ctx context.Context) {
			err := removeTemporaryDir()

			if err != nil {
				println("Error:", err.Error())
			}
		},
		Bind: []interface{}{
			app,
		},
		Linux: &linux.Options{
			Icon:                []byte{},
			WindowIsTranslucent: false,
			Messages:            nil,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyOnDemand,
			ProgramName:         "The Tool",
		},
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func initializeTemporaryDir() error {
	dir := "frontend/src/assets/temp"
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		if os.IsExist(err) {
			return nil
		} else {
			println("Error while creating the temp directory:", err.Error())
			return err
		}
	}
	return nil
}

func removeTemporaryDir() error {
	dir := "frontend/src/assets/temp"
	err := os.RemoveAll(dir)
	if err != nil {
		return err
	}
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

func setMenu(app *App) *menu.Menu {
	AppMenu := menu.NewMenu()
	setFileMenu(app, AppMenu)
	//setGeoTransformMenu(app, AppMenu)
	//setFiltersMenu(app, AppMenu)
	//setMathMorphologyMenu(app, AppMenu)
	//setFeatureExtractionMenu(app, AppMenu)
	return AppMenu
}

// ---------------------------------------------------------------------------------------------------------------------

func setFileMenu(app *App, AppMenu *menu.Menu) {
	// File
	FileMenu := AppMenu.AddSubmenu("File")
	// -- Open image
	FileMenu.AddText("Open", keys.CmdOrCtrl("o"), func(_ *menu.CallbackData) {
		cwd, err := os.Getwd()

		if err != nil {
			println("Error during CWD collection:", err.Error())
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
			println("Error during dialog:", err.Error())
			return
		}

		setOriginPrev(app, path)
	})
	FileMenu.AddSeparator()
	// -- Save image
	FileMenu.AddText("Save", keys.CmdOrCtrl("s"), func(_ *menu.CallbackData) {})
	FileMenu.AddSeparator()
	// -- About
	FileMenu.AddText("About", keys.CmdOrCtrl("f1"), func(_ *menu.CallbackData) {
		runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
			Type:          runtime.InfoDialog,
			Title:         "About",
			Message:       "This is an image manipulation tool written in Go using Wails framework to Run the frontend.",
			DefaultButton: "Back",
		})
	})
	FileMenu.AddSeparator()
	// -- Exit
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
	GeoTransformMenu.AddText("Horizontal mirroring", keys.CmdOrCtrl("h"), func(_ *menu.CallbackData) {})
	GeoTransformMenu.AddSeparator()
	// -- Vertical mirroring
	GeoTransformMenu.AddText("Vertical mirroring", keys.Key("v"), func(_ *menu.CallbackData) {})
	GeoTransformMenu.AddSeparator()
	// -- Resize
	GeoTransformMenu.AddText("Resize", keys.Key("w"), func(_ *menu.CallbackData) {})
}

func setFiltersMenu(app *App, AppMenu *menu.Menu) {
	// Filters
	FiltersMenu := AppMenu.AddSubmenu("Filters")
	// -- Grayscale
	FiltersMenu.AddText("Grayscale", nil, func(_ *menu.CallbackData) {})
	FiltersMenu.AddSeparator()
	// -- Low fade
	FiltersMenu.AddText("Low fade", nil, func(_ *menu.CallbackData) {})
	FiltersMenu.AddSeparator()
	// -- High fade
	FiltersMenu.AddText("High fade", nil, func(_ *menu.CallbackData) {})
	FiltersMenu.AddSeparator()
	// -- Threshold
	FiltersMenu.AddText("Threshold", nil, func(_ *menu.CallbackData) {})
}

func setMathMorphologyMenu(app *App, AppMenu *menu.Menu) {
	// Mathematical morphology
	MathMorphoMenu := AppMenu.AddSubmenu("Mathematical Morphology")
	// -- Dilatation
	MathMorphoMenu.AddText("Dilatation", nil, func(_ *menu.CallbackData) {})
	MathMorphoMenu.AddSeparator()
	// -- Erosion
	MathMorphoMenu.AddText("Erosion", nil, func(_ *menu.CallbackData) {})
	MathMorphoMenu.AddSeparator()
	// -- Opening
	MathMorphoMenu.AddText("Opening", nil, func(_ *menu.CallbackData) {})
	MathMorphoMenu.AddSeparator()
	// -- Closing
	MathMorphoMenu.AddText("Closing", nil, func(_ *menu.CallbackData) {})
}

func setFeatureExtractionMenu(app *App, AppMenu *menu.Menu) {
	// Feature extraction
	AppMenu.AddSubmenu("Feature Extraction")
}

// ---------------------------------------------------------------------------------------------------------------------

func getFileFromMemory(path string) (string, error) {
	fileExt := filepath.Ext(path)
	destPath := "frontend/src/assets/temp/origin" + fileExt

	// Load file
	sourceFile, err := os.Open(path)

	if err != nil {
		println("Error during file opening:", err.Error())
		return "", err
	}
	defer sourceFile.Close()

	// Create new file
	destFile, err := os.Create(destPath)

	if err != nil {
		println("Error during file creation:", err.Error())
		return "", err
	}
	defer destFile.Close()

	// Copy content from file 1 to file 2
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		println("Error during data copying:", err.Error())
		return "", err
	}

	return fileExt, nil
}

func setOriginPrev(app *App, path string) {
	createImage(path, true)
	createImage(path, false)

	runtime.EventsEmit(app.ctx, "set-image", map[string]interface{}{
		"fileExt": filepath.Ext(path),
	})
}

func createImage(originalPath string, origin bool) {
	var path string
	if origin {
		path = "frontend/src/assets/temp/origin" + filepath.Ext(originalPath)
	} else {
		path = "frontend/src/assets/temp/prev" + filepath.Ext(originalPath)
	}

	// Load original file
	originalFile, err := os.Open(originalPath)

	if err != nil {
		println("Error during file opening:", err.Error())
	}
	defer originalFile.Close()

	// Create new file
	destFile, err := os.Create(path)

	if err != nil {
		println("Error during file creation:", err.Error())
		return
	}
	defer destFile.Close()

	// Copy content from file 1 to file 2
	_, err = io.Copy(destFile, originalFile)
	if err != nil {
		println("Error during data copying:", err.Error())
		return
	}
}