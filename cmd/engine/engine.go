package main

import (
	"runtime"

	"github.com/PavlushaSource/Radar/model/core/rnd"
	"github.com/PavlushaSource/Radar/model/engine"
	"github.com/PavlushaSource/Radar/model/geom"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 4)

	// f, err := os.Create("cpu.pprof")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	rndAsync := rnd.NewRndCore()

	g := geom.NewSimpleGeom(2160, 3840, make([]geom.Barrier, 0), geom.EuclideanDistance, rndAsync)
	eng := engine.NewEngine(15, 15, 50000, g, rndAsync, 5)
	// getCh, putCh := eng.Run(context.Background())

	// // start := time.Now()
	// // state := <-getCh
	// // fmt.Println(state.NumCats())
	// // end := time.Now()
	// // fmt.Println("Calculation time:", end.Sub(start))

	for {
		eng.RawRun()
	}

	// fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
	// fmt.Println(runtime.NumCPU())
	// fmt.Println(runtime.NumCgoCall())
	// fmt.Println(runtime.NumGoroutine())

	// for {
	// 	start := time.Now()
	// 	state := <-getCh
	// 	fmt.Println(state.NumCats())
	// 	putCh <- state
	// 	end := time.Now()
	// 	fmt.Println("Calculation time:", end.Sub(start))
	// }

	// start := time.Now()
	// pprof.StartCPUProfile(f)
	// eng.Run(context.Background())
	// pprof.StopCPUProfile()
	// end := time.Now()
	// fmt.Println("Calculate time", end.Sub(start))

}
