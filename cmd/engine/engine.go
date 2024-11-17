package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"

	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/model/geom"
)

func main() {
	f, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	g := geom.NewSimpleGeom(1080, 1920, make([]geom.Barrier, 0), geom.EuclideanDistance)
	eng := engine.NewEngine(15, 30, 50000, g)

	start := time.Now()
	eng.Run()
	end := time.Now()
	fmt.Println("Calculate time", end.Sub(start))

}
