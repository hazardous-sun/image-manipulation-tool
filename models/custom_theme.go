package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type CustomTheme struct {
	data []interface{}
}

func (c *CustomTheme) Name() string {
	return c.data[0].(string)
}

func (c *CustomTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	if n == theme.ColorNameBackground {
		if v == theme.VariantLight {
			return color.RGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0x00}
		}
		return color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}
	}
	return nil
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
