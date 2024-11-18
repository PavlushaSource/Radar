package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func createToolbarFunction(onHomeAction, onThemeAction, onFullScreenAction func()) func() fyne.CanvasObject {
	return func() fyne.CanvasObject {
		topToolbar := widget.Toolbar{
			Items: []widget.ToolbarItem{
				widget.NewToolbarAction(theme.HomeIcon(), func() { onHomeAction() }),
				widget.NewToolbarAction(theme.ColorPaletteIcon(), func() { onThemeAction() }),
				widget.NewToolbarAction(theme.ViewFullScreenIcon(), func() { onFullScreenAction() })},
		}

		return &topToolbar
	}
}
