package utils

import (
	"time"

	"fyne.io/fyne/v2"
)

func AnimateCat(A, B fyne.Position, cat fyne.CanvasObject, iterations int32) {
	// TODO: Add scale logic
	vector := fyne.NewPos((B.X-A.X)/float32(iterations), (B.Y-A.Y)/float32(iterations))

	for i := 0; i < int(iterations); i++ {
		time.Sleep(time.Millisecond)
		cat.Move(cat.Position().Add(vector))
	}
}
