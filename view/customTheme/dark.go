package customTheme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type DarkTheme struct {
	fyne.Theme
}

func NewDarkTheme() fyne.Theme {
	return &DarkTheme{Theme: theme.DefaultTheme()}
}

func (l *DarkTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return l.Theme.Color(name, theme.VariantDark)
}

func (l *DarkTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 14
	}
	return l.Theme.Size(name)
}
