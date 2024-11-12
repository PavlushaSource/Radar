package UI

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/PavlushaSource/Radar/custom_theme"
	"github.com/PavlushaSource/Radar/internal/config"
)

func CreateToolbarFunction(a fyne.App, w1, w2 fyne.Window, Config *config.UIConfig, ctx context.Context) func() fyne.CanvasObject {
	return func() fyne.CanvasObject {
		topToolbar := widget.Toolbar{
			Items: []widget.ToolbarItem{widget.NewToolbarAction(theme.HomeIcon(), func() {
				fmt.Println(*Config)
				if !Config.InMainMenu {

					// TODO develop this later
					//stopRadar(ctx)

					w2.SetFullScreen(false)
					Config.FullScreenMode = false
					w2.Hide()

					w1.CenterOnScreen()
					w1.Resize(Config.WindowSize)
					w1.Show()

					w1.SetMaster()
					Config.InMainMenu = !Config.InMainMenu
				}
			}), widget.NewToolbarAction(theme.ColorPaletteIcon(), func() {
				if Config.LightThemeFlag {
					a.Settings().SetTheme(custom_theme.NewDarkTheme())
				} else {
					a.Settings().SetTheme(custom_theme.NewLightTheme())
				}
				Config.LightThemeFlag = !Config.LightThemeFlag
			}), widget.NewToolbarAction(theme.ViewFullScreenIcon(), func() {
				if !Config.FullScreenMode {
					if Config.InMainMenu {
						w1.SetFullScreen(true)
					} else {
						w2.SetFullScreen(true)
					}
					fmt.Println("PRESS FULL SCREEN", Config.InMainMenu)
				} else {
					if Config.InMainMenu {
						w1.SetFullScreen(false)
					} else {
						w2.SetFullScreen(false)
					}
				}
				Config.FullScreenMode = !Config.FullScreenMode
			})},
		}

		return &topToolbar
	}
}
