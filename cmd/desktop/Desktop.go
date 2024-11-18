package main

import (
	"context"

	"github.com/PavlushaSource/Radar/view"
)

func main() {
	ctx := context.Background()

	app := view.NewApplication()

	app.Run(ctx)
}
