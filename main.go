package main

import (
	"fmt"

	"mlt/lib"
)

func main() {
	a := lib.NewMatrixFloat64([][]float64{
		[]float64{1.01, 2.4, 3.2},
		[]float64{4.16, 5.31, 6.5},
	})
	b := lib.NewMatrixFloat64([][]float64{
		[]float64{1.01, 4.16},
		[]float64{2.4, 5.31},
		[]float64{3.2, 6.5},
	})
	fmt.Println(a.Multi(b))
}
