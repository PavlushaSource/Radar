package main

import (
	"fmt"
	"time"

	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/model/geom"
)

func main() {

	g := geom.NewSimpleGeom(1080, 1920, nil, geom.EuclideanDistance)
	eng := engine.NewEngine(15, 30, 100, g)

	start := time.Now()
	eng.Run()
	end := time.Now()
	fmt.Println("Calculate time", end.Sub(start))

}
