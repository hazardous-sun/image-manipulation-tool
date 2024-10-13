package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type CustomTheme struct {
	name                         string
	ColorNameBackground          *color.RGBA
	ColorNameButton              *color.RGBA
	ColorNameDisabledButton      *color.RGBA
	ColorNameDisabled            *color.RGBA
	ColorNameError               *color.RGBA
	ColorNameFocus               *color.RGBA
	ColorNameForeground          *color.RGBA
	ColorNameForegroundOnError   *color.RGBA
	ColorNameForegroundOnPrimary *color.RGBA
	ColorNameForegroundOnSuccess *color.RGBA
	ColorNameForegroundOnWarning *color.RGBA
	ColorNameHeaderBackground    *color.RGBA
	ColorNameHover               *color.RGBA
	ColorNameHyperlink           *color.RGBA
	ColorNameInputBackground     *color.RGBA
	ColorNameInputBorder         *color.RGBA
	ColorNameMenuBackground      *color.RGBA
	ColorNameOverlayBackground   *color.RGBA
	ColorNamePlaceHolder         *color.RGBA
	ColorNamePressed             *color.RGBA
	ColorNamePrimary             *color.RGBA
	ColorNameScrollBar           *color.RGBA
	ColorNameSelection           *color.RGBA
	ColorNameSeparator           *color.RGBA
	ColorNameShadow              *color.RGBA
	ColorNameSuccess             *color.RGBA
	ColorNameWarning             *color.RGBA
}

func (c *CustomTheme) Name() string {
	return c.name
}

func (c *CustomTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground:
		if c.ColorNameBackground != nil {
			return c.ColorNameBackground
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameButton:
		if c.ColorNameButton != nil {
			return c.ColorNameButton
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameDisabledButton:
		if c.ColorNameDisabledButton != nil {
			return c.ColorNameDisabledButton
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameDisabled:
		if c.ColorNameDisabled != nil {
			return c.ColorNameDisabled
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameError:
		if c.ColorNameError != nil {
			return c.ColorNameError
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameFocus:
		if c.ColorNameFocus != nil {
			return c.ColorNameFocus
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameForeground:
		if c.ColorNameForeground != nil {
			return c.ColorNameForeground
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameForegroundOnError:
		if c.ColorNameForegroundOnError != nil {
			return c.ColorNameForegroundOnError
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameForegroundOnPrimary:
		if c.ColorNameForegroundOnPrimary != nil {
			return c.ColorNameForegroundOnPrimary
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameForegroundOnSuccess:
		if c.ColorNameForegroundOnSuccess != nil {
			return c.ColorNameForegroundOnSuccess
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameForegroundOnWarning:
		if c.ColorNameForegroundOnWarning != nil {
			return c.ColorNameForegroundOnWarning
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameHeaderBackground:
		if c.ColorNameHeaderBackground != nil {
			return c.ColorNameHeaderBackground
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameHover:
		if c.ColorNameHover != nil {
			return c.ColorNameHover
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameHyperlink:
		if c.ColorNameHyperlink != nil {
			return c.ColorNameHyperlink
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameInputBackground:
		if c.ColorNameInputBackground != nil {
			return c.ColorNameInputBackground
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameInputBorder:
		if c.ColorNameInputBorder != nil {
			return c.ColorNameInputBorder
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameMenuBackground:
		if c.ColorNameMenuBackground != nil {
			return c.ColorNameMenuBackground
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameOverlayBackground:
		if c.ColorNameOverlayBackground != nil {
			return c.ColorNameOverlayBackground
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNamePlaceHolder:
		if c.ColorNamePlaceHolder != nil {
			return c.ColorNamePlaceHolder
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNamePressed:
		if c.ColorNamePressed != nil {
			return c.ColorNamePressed
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNamePrimary:
		if c.ColorNamePrimary != nil {
			return c.ColorNamePrimary
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameScrollBar:
		if c.ColorNameScrollBar != nil {
			return c.ColorNameScrollBar
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameSelection:
		if c.ColorNameSelection != nil {
			return c.ColorNameSelection
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameSeparator:
		if c.ColorNameSeparator != nil {
			return c.ColorNameSeparator
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameShadow:
		if c.ColorNameShadow != nil {
			return c.ColorNameShadow
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameSuccess:
		if c.ColorNameSuccess != nil {
			return c.ColorNameSuccess
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameWarning:
		if c.ColorNameWarning != nil {
			return c.ColorNameWarning
		}
		return theme.DefaultTheme().Color(n, v)
	default:
		return theme.DefaultTheme().Color(n, v)
	}
}

func (c *CustomTheme) Font(f fyne.TextStyle) fyne.Resource {
	return theme.Font(
		f,
	)
}

func (c *CustomTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (c *CustomTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}

func NewCustomTheme() *CustomTheme {
	return &CustomTheme{}
}
