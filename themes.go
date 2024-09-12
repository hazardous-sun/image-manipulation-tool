package main

import "github.com/wailsapp/wails/v2/pkg/runtime"

// Theme ---------------------------------------------------------------------------------------------------------------

// Sends a message to the JavaScript listener informing the theme needs to be updated.
func updateTheme(app *App) {
	runtime.EventsEmit(app.ctx, "set-image", map[string]interface{}{})
}
