package main

import (
	"log"

	"mlt/lib"
	"mlt/methods"
)

func main() {
	// train
	var (
		xSet []lib.Vector = []lib.Vector{
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
		ySet []int = []int{1, 1, 1, 1, 1, 1, -1, -1, -1}
	)
	perceptron := methods.NewBasePerceptron(lib.NewVectorFloat64([]float64{-9, 3}), -5, 1)
	perceptron.Train(xSet, ySet)

	// test
	// dst: w=(-1, 1) b=0
	var uk lib.Vector
	for i := 0; i < len(xSet); i++ {

		uk = xSet[i]
		r, e := perceptron.Predict(uk)
		if e != nil {
			log.Printf("error: %s\n", e.Error())
		} else {
			log.Printf("s: %v r: %d\n", uk, r)
		}
	}
}
