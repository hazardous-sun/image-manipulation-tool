package models

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
	"os"
)

// CustomTheme :
// The representation of a custom theme for Fyne. It holds all the data referencing the color for each visual element.
type CustomTheme struct {
	name                         string
	colorNameBackground          *color.RGBA
	colorNameButton              *color.RGBA
	colorNameDisabledButton      *color.RGBA
	colorNameDisabled            *color.RGBA
	colorNameError               *color.RGBA
	colorNameFocus               *color.RGBA
	colorNameForeground          *color.RGBA
	colorNameForegroundOnError   *color.RGBA
	colorNameForegroundOnPrimary *color.RGBA
	colorNameForegroundOnSuccess *color.RGBA
	colorNameForegroundOnWarning *color.RGBA
	colorNameHeaderBackground    *color.RGBA
	colorNameHover               *color.RGBA
	colorNameHyperlink           *color.RGBA
	colorNameInputBackground     *color.RGBA
	colorNameInputBorder         *color.RGBA
	colorNameMenuBackground      *color.RGBA
	colorNameOverlayBackground   *color.RGBA
	colorNamePlaceHolder         *color.RGBA
	colorNamePressed             *color.RGBA
	colorNamePrimary             *color.RGBA
	colorNameScrollBar           *color.RGBA
	colorNameSelection           *color.RGBA
	colorNameSeparator           *color.RGBA
	colorNameShadow              *color.RGBA
	colorNameSuccess             *color.RGBA
	colorNameWarning             *color.RGBA
}

// Name :
// Returns the theme name.
func (t *CustomTheme) Name() string {
	return t.name
}

// Color, Font, Icon and Size are all part of the fyne.Theme interface.
// These methods are called for each new element added to a window, so Fyne knows which color, fonts, icons and size
// to use when drawing the object.

func (t *CustomTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground:
		if t.colorNameBackground != nil {
			return t.colorNameBackground
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameButton:
		if t.colorNameButton != nil {
			return t.colorNameButton
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameDisabledButton:
		if t.colorNameDisabledButton != nil {
			return t.colorNameDisabledButton
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameDisabled:
		if t.colorNameDisabled != nil {
			return t.colorNameDisabled
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameError:
		if t.colorNameError != nil {
			return t.colorNameError
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameFocus:
		if t.colorNameFocus != nil {
			return t.colorNameFocus
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameForeground:
		if t.colorNameForeground != nil {
			return t.colorNameForeground
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameForegroundOnError:
		if t.colorNameForegroundOnError != nil {
			return t.colorNameForegroundOnError
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameForegroundOnPrimary:
		if t.colorNameForegroundOnPrimary != nil {
			return t.colorNameForegroundOnPrimary
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameForegroundOnSuccess:
		if t.colorNameForegroundOnSuccess != nil {
			return t.colorNameForegroundOnSuccess
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameForegroundOnWarning:
		if t.colorNameForegroundOnWarning != nil {
			return t.colorNameForegroundOnWarning
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameHeaderBackground:
		if t.colorNameHeaderBackground != nil {
			return t.colorNameHeaderBackground
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameHover:
		if t.colorNameHover != nil {
			return t.colorNameHover
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameHyperlink:
		if t.colorNameHyperlink != nil {
			return t.colorNameHyperlink
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameInputBackground:
		if t.colorNameInputBackground != nil {
			return t.colorNameInputBackground
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameInputBorder:
		if t.colorNameInputBorder != nil {
			return t.colorNameInputBorder
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameMenuBackground:
		if t.colorNameMenuBackground != nil {
			return t.colorNameMenuBackground
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameOverlayBackground:
		if t.colorNameOverlayBackground != nil {
			return t.colorNameOverlayBackground
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNamePlaceHolder:
		if t.colorNamePlaceHolder != nil {
			return t.colorNamePlaceHolder
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNamePressed:
		if t.colorNamePressed != nil {
			return t.colorNamePressed
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNamePrimary:
		if t.colorNamePrimary != nil {
			return t.colorNamePrimary
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameScrollBar:
		if t.colorNameScrollBar != nil {
			return t.colorNameScrollBar
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameSelection:
		if t.colorNameSelection != nil {
			return t.colorNameSelection
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameSeparator:
		if t.colorNameSeparator != nil {
			return t.colorNameSeparator
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameShadow:
		if t.colorNameShadow != nil {
			return t.colorNameShadow
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameSuccess:
		if t.colorNameSuccess != nil {
			return t.colorNameSuccess
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameWarning:
		if t.colorNameWarning != nil {
			return t.colorNameWarning
		}
		return theme.DefaultTheme().Color(n, v)
	default:
		return theme.DefaultTheme().Color(n, v)
	}
}

func (t *CustomTheme) Font(s fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(s)
}

func (t *CustomTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (t *CustomTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}

// UnmarshalJSON is an implementation of an interface that allows the values from a theme to be loaded from a JSON file
// into an instance of CustomTheme.

func (t *CustomTheme) UnmarshalJSON(data []byte) error {
	type RawCustomTheme struct {
		Name                string
		ColorNameBackground struct {
			R, G, B, A uint8
		}
		ColorNameButton struct {
			R, G, B, A uint8
		}
		ColorNameDisabledButton struct {
			R, G, B, A uint8
		}
		ColorNameDisabled struct {
			R, G, B, A uint8
		}
		ColorNameError struct {
			R, G, B, A uint8
		}
		ColorNameFocus struct {
			R, G, B, A uint8
		}
		ColorNameForeground struct {
			R, G, B, A uint8
		}
		ColorNameForegroundOnError struct {
			R, G, B, A uint8
		}
		ColorNameForegroundOnPrimary struct {
			R, G, B, A uint8
		}
		ColorNameForegroundOnSuccess struct {
			R, G, B, A uint8
		}
		ColorNameForegroundOnWarning struct {
			R, G, B, A uint8
		}
		ColorNameHeaderBackground struct {
			R, G, B, A uint8
		}
		ColorNameHover struct {
			R, G, B, A uint8
		}
		ColorNameHyperlink struct {
			R, G, B, A uint8
		}
		ColorNameInputBackground struct {
			R, G, B, A uint8
		}
		ColorNameInputBorder struct {
			R, G, B, A uint8
		}
		ColorNameMenuBackground struct {
			R, G, B, A uint8
		}
		ColorNameOverlayBackground struct {
			R, G, B, A uint8
		}
		ColorNamePlaceHolder struct {
			R, G, B, A uint8
		}
		ColorNamePressed struct {
			R, G, B, A uint8
		}
		ColorNamePrimary struct {
			R, G, B, A uint8
		}
		ColorNameScrollBar struct {
			R, G, B, A uint8
		}
		ColorNameSelection struct {
			R, G, B, A uint8
		}
		ColorNameSeparator struct {
			R, G, B, A uint8
		}
		ColorNameShadow struct {
			R, G, B, A uint8
		}
		ColorNameSuccess struct {
			R, G, B, A uint8
		}
		ColorNameWarning struct {
			R, G, B, A uint8
		}
	}

	var raw RawCustomTheme

	err := json.Unmarshal(data, &raw)

	if err != nil {
		return err
	}

	t.name = raw.Name
	t.colorNameBackground = &color.RGBA{
		R: raw.ColorNameBackground.R,
		G: raw.ColorNameBackground.G,
		B: raw.ColorNameBackground.B,
		A: raw.ColorNameBackground.A,
	}
	t.colorNameButton = &color.RGBA{
		R: raw.ColorNameButton.R,
		G: raw.ColorNameButton.G,
		B: raw.ColorNameButton.B,
		A: raw.ColorNameButton.A,
	}
	t.colorNameDisabledButton = &color.RGBA{
		R: raw.ColorNameDisabledButton.R,
		G: raw.ColorNameDisabledButton.G,
		B: raw.ColorNameDisabledButton.B,
		A: raw.ColorNameDisabledButton.A,
	}
	t.colorNameDisabled = &color.RGBA{
		R: raw.ColorNameDisabled.R,
		G: raw.ColorNameDisabled.G,
		B: raw.ColorNameDisabled.B,
		A: raw.ColorNameDisabled.A,
	}
	t.colorNameError = &color.RGBA{
		R: raw.ColorNameError.R,
		G: raw.ColorNameError.G,
		B: raw.ColorNameError.B,
		A: raw.ColorNameError.A,
	}
	t.colorNameFocus = &color.RGBA{
		R: raw.ColorNameFocus.R,
		G: raw.ColorNameFocus.G,
		B: raw.ColorNameFocus.B,
		A: raw.ColorNameFocus.A,
	}
	t.colorNameForeground = &color.RGBA{
		R: raw.ColorNameForeground.R,
		G: raw.ColorNameForeground.G,
		B: raw.ColorNameForeground.B,
		A: raw.ColorNameForeground.A,
	}
	t.colorNameForegroundOnError = &color.RGBA{
		R: raw.ColorNameForegroundOnError.R,
		G: raw.ColorNameForegroundOnError.G,
		B: raw.ColorNameForegroundOnError.B,
		A: raw.ColorNameForegroundOnError.A,
	}
	t.colorNameForegroundOnPrimary = &color.RGBA{
		R: raw.ColorNameForegroundOnPrimary.R,
		G: raw.ColorNameForegroundOnPrimary.G,
		B: raw.ColorNameForegroundOnPrimary.B,
		A: raw.ColorNameForegroundOnPrimary.A,
	}
	t.colorNameForegroundOnSuccess = &color.RGBA{
		R: raw.ColorNameForegroundOnSuccess.R,
		G: raw.ColorNameForegroundOnSuccess.G,
		B: raw.ColorNameForegroundOnSuccess.B,
		A: raw.ColorNameForegroundOnSuccess.A,
	}
	t.colorNameForegroundOnWarning = &color.RGBA{
		R: raw.ColorNameForegroundOnWarning.R,
		G: raw.ColorNameForegroundOnWarning.G,
		B: raw.ColorNameForegroundOnWarning.B,
		A: raw.ColorNameForegroundOnWarning.A,
	}
	t.colorNameHeaderBackground = &color.RGBA{
		R: raw.ColorNameHeaderBackground.R,
		G: raw.ColorNameHeaderBackground.G,
		B: raw.ColorNameHeaderBackground.B,
		A: raw.ColorNameHeaderBackground.A,
	}
	t.colorNameHover = &color.RGBA{
		R: raw.ColorNameHover.R,
		G: raw.ColorNameHover.G,
		B: raw.ColorNameHover.B,
		A: raw.ColorNameHover.A,
	}
	t.colorNameHyperlink = &color.RGBA{
		R: raw.ColorNameHyperlink.R,
		G: raw.ColorNameHyperlink.G,
		B: raw.ColorNameHyperlink.B,
		A: raw.ColorNameHyperlink.A,
	}
	t.colorNameInputBackground = &color.RGBA{
		R: raw.ColorNameInputBackground.R,
		G: raw.ColorNameInputBackground.G,
		B: raw.ColorNameInputBackground.B,
		A: raw.ColorNameInputBackground.A,
	}
	t.colorNameInputBorder = &color.RGBA{
		R: raw.ColorNameInputBorder.R,
		G: raw.ColorNameInputBorder.G,
		B: raw.ColorNameInputBorder.B,
		A: raw.ColorNameInputBorder.A,
	}
	t.colorNameMenuBackground = &color.RGBA{
		R: raw.ColorNameMenuBackground.R,
		G: raw.ColorNameMenuBackground.G,
		B: raw.ColorNameMenuBackground.B,
		A: raw.ColorNameMenuBackground.A,
	}
	t.colorNameOverlayBackground = &color.RGBA{
		R: raw.ColorNameOverlayBackground.R,
		G: raw.ColorNameOverlayBackground.G,
		B: raw.ColorNameOverlayBackground.B,
		A: raw.ColorNameOverlayBackground.A,
	}
	t.colorNamePlaceHolder = &color.RGBA{
		R: raw.ColorNamePlaceHolder.R,
		G: raw.ColorNamePlaceHolder.G,
		B: raw.ColorNamePlaceHolder.B,
		A: raw.ColorNamePlaceHolder.A,
	}
	t.colorNamePressed = &color.RGBA{
		R: raw.ColorNamePressed.R,
		G: raw.ColorNamePressed.G,
		B: raw.ColorNamePressed.B,
		A: raw.ColorNamePressed.A,
	}
	t.colorNamePrimary = &color.RGBA{
		R: raw.ColorNamePrimary.R,
		G: raw.ColorNamePrimary.G,
		B: raw.ColorNamePrimary.B,
		A: raw.ColorNamePrimary.A,
	}
	t.colorNameScrollBar = &color.RGBA{
		R: raw.ColorNameScrollBar.R,
		G: raw.ColorNameScrollBar.G,
		B: raw.ColorNameScrollBar.B,
		A: raw.ColorNameScrollBar.A,
	}
	t.colorNameSelection = &color.RGBA{
		R: raw.ColorNameSelection.R,
		G: raw.ColorNameSelection.G,
		B: raw.ColorNameSelection.B,
		A: raw.ColorNameSelection.A,
	}
	t.colorNameSeparator = &color.RGBA{
		R: raw.ColorNameSeparator.R,
		G: raw.ColorNameSeparator.G,
		B: raw.ColorNameSeparator.B,
		A: raw.ColorNameSeparator.A,
	}
	t.colorNameShadow = &color.RGBA{
		R: raw.ColorNameShadow.R,
		G: raw.ColorNameShadow.G,
		B: raw.ColorNameShadow.B,
		A: raw.ColorNameShadow.A,
	}
	t.colorNameSuccess = &color.RGBA{
		R: raw.ColorNameSuccess.R,
		G: raw.ColorNameSuccess.G,
		B: raw.ColorNameSuccess.B,
		A: raw.ColorNameSuccess.A,
	}
	t.colorNameWarning = &color.RGBA{
		R: raw.ColorNameWarning.R,
		G: raw.ColorNameWarning.G,
		B: raw.ColorNameWarning.B,
		A: raw.ColorNameWarning.A,
	}

	return nil
}

// NewCustomTheme :
// Returns an instance of CustomTheme based on a JSON file provided through the "path" parameter.
func NewCustomTheme(path string) (*CustomTheme, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing file:", err)
			return
		}
	}()

	var customTheme CustomTheme
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&customTheme)

	if err != nil {
		return nil, err
	}

	return &customTheme, nil
}
