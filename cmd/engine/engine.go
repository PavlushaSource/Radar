package main

import (
	"context"
	"fmt"
  "os"
	"time"
  
  "github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/model/geom"
)

func main() {
	f, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rndAsync := rnd.NewRndCore()

	g := geom.NewSimpleGeom(1080, 1920, make([]geom.Barrier, 0), geom.EuclideanDistance, rndAsync)
	eng := engine.NewEngine(1, 2, 50000, g, rndAsync, 5)
	getCh, putCh := eng.Run(context.Background())

	// start := time.Now()
	// state := <-getCh
	// fmt.Println(state.NumCats())
	// end := time.Now()
	// fmt.Println("Calculation time:", end.Sub(start))

	for {
		start := time.Now()
		state := <-getCh
		fmt.Println(state.NumCats())
		putCh <- state
		end := time.Now()
		fmt.Println("Calculation time:", end.Sub(start))
	}

	// start := time.Now()
	// pprof.StartCPUProfile(f)
	// eng.Run(context.Background())
	// pprof.StopCPUProfile()
	// end := time.Now()
	// fmt.Println("Calculate time", end.Sub(start))

}
