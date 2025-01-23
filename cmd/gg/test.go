package main

import (
	"github.com/PavlushaSource/Radar/view"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {

	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Ebiten UI - TextInput with Labels")

	app := &view.Application{InMainMenu: true}

	app.Menu = view.NewUI(app)

	err := ebiten.RunGame(app)
	if err != nil {
		log.Println(err)
	}
}
