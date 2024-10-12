package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type ThemeSettings struct {
	currentTheme fyne.Theme
	themesList   []TestTheme
}

func (s *ThemeSettings) Theme() fyne.Theme {
	return s.currentTheme
}

func (s *ThemeSettings) Themes() []TestTheme {
	return s.themesList
}

func (s *ThemeSettings) AddTheme(theme TestTheme) {
	s.themesList = append(s.themesList, theme)
}

func (s *ThemeSettings) NewTheme() {

}

func NewThemeSettings() *ThemeSettings {
	t := ThemeSettings{
		currentTheme: theme.DefaultTheme(),
		themesList:   []TestTheme{},
	}
	t.AddTheme(TestTheme{
		"Dark",
		theme.DarkTheme(),
	})
	t.AddTheme(TestTheme{
		"Light",
		theme.LightTheme(),
	})

	return &t
}

type TestTheme struct {
	Name  string
	Theme fyne.Theme
}
