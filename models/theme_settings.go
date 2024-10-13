package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type ThemeSettings struct {
	currentTheme fyne.Theme
	themesList   []CustomTheme
}

func (s *ThemeSettings) Theme() fyne.Theme {
	return s.currentTheme
}

func (s *ThemeSettings) Themes() []CustomTheme {
	return s.themesList
}

func (s *ThemeSettings) AddTheme(theme CustomTheme) {
	s.themesList = append(s.themesList, theme)
}

func (s *ThemeSettings) NewTheme() {

}

func NewThemeSettings() *ThemeSettings {
	t := ThemeSettings{
		currentTheme: theme.DefaultTheme(),
		themesList:   []CustomTheme{},
	}

	return &t
}
