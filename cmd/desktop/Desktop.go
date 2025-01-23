package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"

	"github.com/PavlushaSource/Radar/view"
)

func main() {

	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Menu")

	app := view.NewApplication()

	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
