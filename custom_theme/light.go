package custom_theme

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type LightTheme struct {
	fyne.Theme
}

func NewLightTheme() fyne.Theme {
	return &LightTheme{Theme: theme.DefaultTheme()}
}

func (l *LightTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return l.Theme.Color(name, theme.VariantLight)
}

func (l *LightTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 14
	}
	return l.Theme.Size(name)
}
