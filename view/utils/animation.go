package utils

import (
	"time"

	"fyne.io/fyne/v2"
)

func AnimateDog(A, B fyne.Position, dog fyne.CanvasObject, iterations int32) {
	// TODO: Add scale logic
	vector := fyne.NewPos((B.X-A.X)/float32(iterations)*16, (B.Y-A.Y)/float32(iterations)*16)

	for i := 0; i < int(iterations)/16; i++ {
		time.Sleep(time.Millisecond * 16)
		dog.Move(dog.Position().Add(vector))
	}
}
