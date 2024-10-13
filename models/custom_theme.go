package models

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
	"os"
)

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

func (t *CustomTheme) Name() string {
	return t.name
}

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

func (t *CustomTheme) UnmarshalJSON(data []byte) error {
	type RawCustomTheme struct {
		Name                string
		colorNameBackground struct {
			R, G, B, A uint8
		}
		colorNameButton struct {
			R, G, B, A uint8
		}
		colorNameDisabledButton struct {
			R, G, B, A uint8
		}
		colorNameDisabled struct {
			R, G, B, A uint8
		}
		colorNameError struct {
			R, G, B, A uint8
		}
		colorNameFocus struct {
			R, G, B, A uint8
		}
		colorNameForeground struct {
			R, G, B, A uint8
		}
		colorNameForegroundOnError struct {
			R, G, B, A uint8
		}
		colorNameForegroundOnPrimary struct {
			R, G, B, A uint8
		}
		colorNameForegroundOnSuccess struct {
			R, G, B, A uint8
		}
		colorNameForegroundOnWarning struct {
			R, G, B, A uint8
		}
		colorNameHeaderBackground struct {
			R, G, B, A uint8
		}
		colorNameHover struct {
			R, G, B, A uint8
		}
		colorNameHyperlink struct {
			R, G, B, A uint8
		}
		colorNameInputBackground struct {
			R, G, B, A uint8
		}
		colorNameInputBorder struct {
			R, G, B, A uint8
		}
		colorNameMenuBackground struct {
			R, G, B, A uint8
		}
		colorNameOverlayBackground struct {
			R, G, B, A uint8
		}
		colorNamePlaceHolder struct {
			R, G, B, A uint8
		}
		colorNamePressed struct {
			R, G, B, A uint8
		}
		colorNamePrimary struct {
			R, G, B, A uint8
		}
		colorNameScrollBar struct {
			R, G, B, A uint8
		}
		colorNameSelection struct {
			R, G, B, A uint8
		}
		colorNameSeparator struct {
			R, G, B, A uint8
		}
		colorNameShadow struct {
			R, G, B, A uint8
		}
		colorNameSuccess struct {
			R, G, B, A uint8
		}
		colorNameWarning struct {
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
		R: raw.colorNameBackground.R,
		G: raw.colorNameBackground.G,
		B: raw.colorNameBackground.B,
		A: raw.colorNameBackground.A,
	}
	t.colorNameBackground = &color.RGBA{
		R: raw.colorNameBackground.R,
		G: raw.colorNameBackground.G,
		B: raw.colorNameBackground.B,
		A: raw.colorNameBackground.A,
	}
	t.colorNameButton = &color.RGBA{
		R: raw.colorNameButton.R,
		G: raw.colorNameButton.G,
		B: raw.colorNameButton.B,
		A: raw.colorNameButton.A,
	}
	t.colorNameDisabledButton = &color.RGBA{
		R: raw.colorNameDisabledButton.R,
		G: raw.colorNameDisabledButton.G,
		B: raw.colorNameDisabledButton.B,
		A: raw.colorNameDisabledButton.A,
	}
	t.colorNameDisabled = &color.RGBA{
		R: raw.colorNameDisabled.R,
		G: raw.colorNameDisabled.G,
		B: raw.colorNameDisabled.B,
		A: raw.colorNameDisabled.A,
	}
	t.colorNameError = &color.RGBA{
		R: raw.colorNameError.R,
		G: raw.colorNameError.G,
		B: raw.colorNameError.B,
		A: raw.colorNameError.A,
	}
	t.colorNameFocus = &color.RGBA{
		R: raw.colorNameFocus.R,
		G: raw.colorNameFocus.G,
		B: raw.colorNameFocus.B,
		A: raw.colorNameFocus.A,
	}
	t.colorNameForeground = &color.RGBA{
		R: raw.colorNameForeground.R,
		G: raw.colorNameForeground.G,
		B: raw.colorNameForeground.B,
		A: raw.colorNameForeground.A,
	}
	t.colorNameForegroundOnError = &color.RGBA{
		R: raw.colorNameForegroundOnError.R,
		G: raw.colorNameForegroundOnError.G,
		B: raw.colorNameForegroundOnError.B,
		A: raw.colorNameForegroundOnError.A,
	}
	t.colorNameForegroundOnPrimary = &color.RGBA{
		R: raw.colorNameForegroundOnPrimary.R,
		G: raw.colorNameForegroundOnPrimary.G,
		B: raw.colorNameForegroundOnPrimary.B,
		A: raw.colorNameForegroundOnPrimary.A,
	}
	t.colorNameForegroundOnSuccess = &color.RGBA{
		R: raw.colorNameForegroundOnSuccess.R,
		G: raw.colorNameForegroundOnSuccess.G,
		B: raw.colorNameForegroundOnSuccess.B,
		A: raw.colorNameForegroundOnSuccess.A,
	}
	t.colorNameForegroundOnWarning = &color.RGBA{
		R: raw.colorNameForegroundOnWarning.R,
		G: raw.colorNameForegroundOnWarning.G,
		B: raw.colorNameForegroundOnWarning.B,
		A: raw.colorNameForegroundOnWarning.A,
	}
	t.colorNameHeaderBackground = &color.RGBA{
		R: raw.colorNameHeaderBackground.R,
		G: raw.colorNameHeaderBackground.G,
		B: raw.colorNameHeaderBackground.B,
		A: raw.colorNameHeaderBackground.A,
	}
	t.colorNameHover = &color.RGBA{
		R: raw.colorNameHover.R,
		G: raw.colorNameHover.G,
		B: raw.colorNameHover.B,
		A: raw.colorNameHover.A,
	}
	t.colorNameHyperlink = &color.RGBA{
		R: raw.colorNameHyperlink.R,
		G: raw.colorNameHyperlink.G,
		B: raw.colorNameHyperlink.B,
		A: raw.colorNameHyperlink.A,
	}
	t.colorNameInputBackground = &color.RGBA{
		R: raw.colorNameInputBackground.R,
		G: raw.colorNameInputBackground.G,
		B: raw.colorNameInputBackground.B,
		A: raw.colorNameInputBackground.A,
	}
	t.colorNameInputBorder = &color.RGBA{
		R: raw.colorNameInputBorder.R,
		G: raw.colorNameInputBorder.G,
		B: raw.colorNameInputBorder.B,
		A: raw.colorNameInputBorder.A,
	}
	t.colorNameMenuBackground = &color.RGBA{
		R: raw.colorNameMenuBackground.R,
		G: raw.colorNameMenuBackground.G,
		B: raw.colorNameMenuBackground.B,
		A: raw.colorNameMenuBackground.A,
	}
	t.colorNameOverlayBackground = &color.RGBA{
		R: raw.colorNameOverlayBackground.R,
		G: raw.colorNameOverlayBackground.G,
		B: raw.colorNameOverlayBackground.B,
		A: raw.colorNameOverlayBackground.A,
	}
	t.colorNamePlaceHolder = &color.RGBA{
		R: raw.colorNamePlaceHolder.R,
		G: raw.colorNamePlaceHolder.G,
		B: raw.colorNamePlaceHolder.B,
		A: raw.colorNamePlaceHolder.A,
	}
	t.colorNamePressed = &color.RGBA{
		R: raw.colorNamePressed.R,
		G: raw.colorNamePressed.G,
		B: raw.colorNamePressed.B,
		A: raw.colorNamePressed.A,
	}
	t.colorNamePrimary = &color.RGBA{
		R: raw.colorNamePrimary.R,
		G: raw.colorNamePrimary.G,
		B: raw.colorNamePrimary.B,
		A: raw.colorNamePrimary.A,
	}
	t.colorNameScrollBar = &color.RGBA{
		R: raw.colorNameScrollBar.R,
		G: raw.colorNameScrollBar.G,
		B: raw.colorNameScrollBar.B,
		A: raw.colorNameScrollBar.A,
	}
	t.colorNameSelection = &color.RGBA{
		R: raw.colorNameSelection.R,
		G: raw.colorNameSelection.G,
		B: raw.colorNameSelection.B,
		A: raw.colorNameSelection.A,
	}
	t.colorNameSeparator = &color.RGBA{
		R: raw.colorNameSeparator.R,
		G: raw.colorNameSeparator.G,
		B: raw.colorNameSeparator.B,
		A: raw.colorNameSeparator.A,
	}
	t.colorNameShadow = &color.RGBA{
		R: raw.colorNameShadow.R,
		G: raw.colorNameShadow.G,
		B: raw.colorNameShadow.B,
		A: raw.colorNameShadow.A,
	}
	t.colorNameSuccess = &color.RGBA{
		R: raw.colorNameSuccess.R,
		G: raw.colorNameSuccess.G,
		B: raw.colorNameSuccess.B,
		A: raw.colorNameSuccess.A,
	}
	t.colorNameWarning = &color.RGBA{
		R: raw.colorNameWarning.R,
		G: raw.colorNameWarning.G,
		B: raw.colorNameWarning.B,
		A: raw.colorNameWarning.A,
	}

	return nil
}

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
