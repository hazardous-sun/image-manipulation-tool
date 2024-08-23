package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io"
	"os"
	"path/filepath"
	goRuntime "runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	err := initializeTemporaryDir()

	if err != nil { // TODO implement a way to delete the directory with a graceful shutdown
		if os.IsExist(err) {
			println(err.Error())
		} else {
			println(err.Error())
			return
		}
	}

	// Create an instance of the app structure
	app := NewApp()

	// Set menu
	AppMenu := setMenuBar(app)

	// Create application with options
	err = wails.Run(&options.App{
		Title:     "Image Manipulation Tool",
		Width:     1366,
		MinWidth:  800,
		Height:    768,
		MinHeight: 1000,
		Menu:      AppMenu,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 5, G: 255, B: 255, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func initializeTemporaryDir() error {
	dir := "frontend/src/assets/temp"
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		println("Error while creating the temp directory:", err.Error())
		return err
	}
	return nil
}

func setMenuBar(app *App) *menu.Menu {
	AppMenu := menu.NewMenu()
	setFileMenu(app, AppMenu)
	//setGeoTransformMenu(app, AppMenu)
	//setFiltersMenu(app, AppMenu)
	//setMathMorphologyMenu(app, AppMenu)
	//setFeatureExtractionMenu(app, AppMenu)
	return AppMenu
}

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

		path, err = duplicateFile(path)

		if err != nil {
			println("Error during file duplication:", err.Error())
			return
		}

		runtime.EventsEmit(app.ctx, "set-image", map[string]interface{}{
			"image": path,
		})

		println(cwd)
		println(path)
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
			Message:       "This is an image manipulation tool written in Go using Wails framework to build the frontend.",
			DefaultButton: "Back",
		})
	})
	FileMenu.AddSeparator()
	// -- Exit
	FileMenu.AddText("Exit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})
}

func duplicateFile(path string) (string, error) {
	fileName := filepath.Base(path)
	destPath := "frontend/src/assets/temp/" + fileName
	sourceFile, err := os.Open(path)

	if err != nil {
		println("Error during file opening:", err.Error())
		return "", err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(destPath)

	if err != nil {
		println("Error during file creation:", err.Error())
		return "", err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		println("Error during data copying:", err.Error())
		return "", err
	}

	return fileName, nil
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
