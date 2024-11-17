package main

import (
	"context"

	"github.com/PavlushaSource/Radar/view/UI"
)

func main() {
	ctx := context.Background()

	app := UI.NewApplication()

	app.Run(ctx)
}
