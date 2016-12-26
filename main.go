package main

import (
	"fmt"

	"mlt/lib"
)

func main() {
	a := lib.NewMatrixInt64([][]int64{
		[]int64{1, 2, 3},
		[]int64{4, 5, 6},
	})
	b := lib.NewMatrixInt64([][]int64{
		[]int64{1, 4},
		[]int64{2, 5},
		[]int64{3, 6},
	})
	fmt.Println(a.Multi(b))
}
