package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// ThemeSettings :
// A structure that holds the current theme of the application, as well as a collection of other themes that were
// loaded into the memory.
type ThemeSettings struct {
	currentTheme fyne.Theme
	themesList   []CustomTheme
}

// Theme :
// Returns the current theme of the app.
func (s *ThemeSettings) Theme() fyne.Theme {
	return s.currentTheme
}

// Themes :
// Returns all the themes already loaded.
func (s *ThemeSettings) Themes() []CustomTheme {
	return s.themesList
}

// AddTheme :
// Adds a new theme to the memory.
func (s *ThemeSettings) AddTheme(theme CustomTheme) {
	s.themesList = append(s.themesList, theme)
}

// NewThemeSettings :
// Returns a reference to an instance of ThemeSettings with ThemeSettings.currentTheme loaded as theme.DefaultTheme.
func NewThemeSettings() *ThemeSettings {
	t := ThemeSettings{
		currentTheme: theme.DefaultTheme(),
		themesList:   []CustomTheme{},
	}

	return &t
}
