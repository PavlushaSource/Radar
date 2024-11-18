package view

import (
	"fyne.io/fyne/v2"
	"time"
)

func AnimateCat(A, B fyne.Position, cat fyne.CanvasObject, iterations float32) {
	// TODO: Add scale logic
	vector := fyne.NewPos((B.X-A.X)/iterations, (B.Y-A.Y)/iterations)

	for i := 0; i < int(iterations); i++ {
		time.Sleep(time.Millisecond)
		cat.Move(cat.Position().Add(vector))
	}
}
