package main

import (
	"log"

	"mlt/lib"
	"mlt/methods"
)

func main() {
	// train
	var (
		set []lib.Vector = []lib.Vector{
			lib.NewVectorFloat64([]float64{-2.0, -0.9}), // 1
			lib.NewVectorFloat64([]float64{-1.5, -2.5}), // 1
			lib.NewVectorFloat64([]float64{-1.0, -0.5}), // 1
			lib.NewVectorFloat64([]float64{-0.5, -0.5}), // 1
			lib.NewVectorFloat64([]float64{0.1, -0.5}),  // 1
			lib.NewVectorFloat64([]float64{0.5, -0.6}),  // 1
			lib.NewVectorFloat64([]float64{1.1, 1.2}),   // -1
			lib.NewVectorFloat64([]float64{1.5, 6.5}),   // -1
			lib.NewVectorFloat64([]float64{2.2, 4.3}),   // -1
		}
	)

	// test
}
