package UI

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/PavlushaSource/Radar/internal/config"
	"github.com/PavlushaSource/Radar/internal/runner"
	"image/color"
	"time"
)

func CreateContentRunWindow(w1, w2 fyne.Window, backConf *config.BackendConfig, frontConf *config.UIConfig, ctx context.Context) fyne.CanvasObject {
	rect := canvas.NewRectangle(color.Transparent)

	r := runner.NewRadar()
	r.Bg = rect
	r.CatSize = frontConf.CatSize
	r.UpdateTime = backConf.UpdateTime

	//backEnd := GenerateBackendCats(50000)
	backEnd := []runner.CatBackend{{X: 0, Y: 0, Color: runner.Red}, {X: 100, Y: 100, Color: runner.Purple}}
	r.Cats = r.CreateCats(backEnd)

	go func() {
		time.Sleep(1 * time.Second)
		r.Run(ctx)
	}()

	cnt := container.NewStack(rect, r)
	return cnt
}
