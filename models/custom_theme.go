package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
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

func (c *CustomTheme) Name() string {
	return c.name
}

func (c *CustomTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground:
		if c.colorNameBackground != nil {
			return c.colorNameBackground
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameButton:
		if c.colorNameButton != nil {
			return c.colorNameButton
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameDisabledButton:
		if c.colorNameDisabledButton != nil {
			return c.colorNameDisabledButton
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameDisabled:
		if c.colorNameDisabled != nil {
			return c.colorNameDisabled
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameError:
		if c.colorNameError != nil {
			return c.colorNameError
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameFocus:
		if c.colorNameFocus != nil {
			return c.colorNameFocus
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameForeground:
		if c.colorNameForeground != nil {
			return c.colorNameForeground
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameForegroundOnError:
		if c.colorNameForegroundOnError != nil {
			return c.colorNameForegroundOnError
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameForegroundOnPrimary:
		if c.colorNameForegroundOnPrimary != nil {
			return c.colorNameForegroundOnPrimary
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameForegroundOnSuccess:
		if c.colorNameForegroundOnSuccess != nil {
			return c.colorNameForegroundOnSuccess
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameForegroundOnWarning:
		if c.colorNameForegroundOnWarning != nil {
			return c.colorNameForegroundOnWarning
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameHeaderBackground:
		if c.colorNameHeaderBackground != nil {
			return c.colorNameHeaderBackground
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameHover:
		if c.colorNameHover != nil {
			return c.colorNameHover
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameHyperlink:
		if c.colorNameHyperlink != nil {
			return c.colorNameHyperlink
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameInputBackground:
		if c.colorNameInputBackground != nil {
			return c.colorNameInputBackground
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameInputBorder:
		if c.colorNameInputBorder != nil {
			return c.colorNameInputBorder
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameMenuBackground:
		if c.colorNameMenuBackground != nil {
			return c.colorNameMenuBackground
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameOverlayBackground:
		if c.colorNameOverlayBackground != nil {
			return c.colorNameOverlayBackground
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNamePlaceHolder:
		if c.colorNamePlaceHolder != nil {
			return c.colorNamePlaceHolder
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNamePressed:
		if c.colorNamePressed != nil {
			return c.colorNamePressed
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNamePrimary:
		if c.colorNamePrimary != nil {
			return c.colorNamePrimary
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameScrollBar:
		if c.colorNameScrollBar != nil {
			return c.colorNameScrollBar
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameSelection:
		if c.colorNameSelection != nil {
			return c.colorNameSelection
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameSeparator:
		if c.colorNameSeparator != nil {
			return c.colorNameSeparator
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameShadow:
		if c.colorNameShadow != nil {
			return c.colorNameShadow
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameSuccess:
		if c.colorNameSuccess != nil {
			return c.colorNameSuccess
		}
		return theme.DefaultTheme().Color(n, v)
	case theme.ColorNameWarning:
		if c.colorNameWarning != nil {
			return c.colorNameWarning
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

func NewCustomTheme(options []ThemeLoader) *CustomTheme {
	newCustom := CustomTheme{}
	for _, v := range options {
		colorName := v.id
		switch colorName {
		case theme.ColorNameBackground:
			newCustom.colorNameBackground = &v.rgb
		case theme.ColorNameButton:
			newCustom.colorNameButton = &v.rgb
		case theme.ColorNameDisabledButton:
			newCustom.colorNameDisabledButton = &v.rgb
		case theme.ColorNameDisabled:
			newCustom.colorNameDisabled = &v.rgb
		case theme.ColorNameError:
			newCustom.colorNameError = &v.rgb
		case theme.ColorNameFocus:
			newCustom.colorNameFocus = &v.rgb
		case theme.ColorNameForeground:
			newCustom.colorNameForeground = &v.rgb
		case theme.ColorNameForegroundOnError:
			newCustom.colorNameForegroundOnError = &v.rgb
		case theme.ColorNameForegroundOnPrimary:
			newCustom.colorNameForegroundOnPrimary = &v.rgb
		case theme.ColorNameForegroundOnSuccess:
			newCustom.colorNameForegroundOnSuccess = &v.rgb
		case theme.ColorNameForegroundOnWarning:
			newCustom.colorNameForegroundOnWarning = &v.rgb
		case theme.ColorNameHeaderBackground:
			newCustom.colorNameHeaderBackground = &v.rgb
		case theme.ColorNameHover:
			newCustom.colorNameHover = &v.rgb
		case theme.ColorNameHyperlink:
			newCustom.colorNameHyperlink = &v.rgb
		case theme.ColorNameInputBackground:
			newCustom.colorNameInputBackground = &v.rgb
		case theme.ColorNameInputBorder:
			newCustom.colorNameInputBorder = &v.rgb
		case theme.ColorNameMenuBackground:
			newCustom.colorNameMenuBackground = &v.rgb
		case theme.ColorNameOverlayBackground:
			newCustom.colorNameOverlayBackground = &v.rgb
		case theme.ColorNamePlaceHolder:
			newCustom.colorNamePlaceHolder = &v.rgb
		case theme.ColorNamePressed:
			newCustom.colorNamePressed = &v.rgb
		case theme.ColorNamePrimary:
			newCustom.colorNamePrimary = &v.rgb
		case theme.ColorNameScrollBar:
			newCustom.colorNameScrollBar = &v.rgb
		case theme.ColorNameSelection:
			newCustom.colorNameSelection = &v.rgb
		case theme.ColorNameSeparator:
			newCustom.colorNameSeparator = &v.rgb
		case theme.ColorNameShadow:
			newCustom.colorNameShadow = &v.rgb
		case theme.ColorNameSuccess:
			newCustom.colorNameSuccess = &v.rgb
		case theme.ColorNameWarning:
			newCustom.colorNameWarning = &v.rgb

		default:
			continue
		}
	}
	return &newCustom
}
