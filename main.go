package main

import (
	"fmt"

	"mlt/lib"
)

func main() {
	a := lib.NewMatrixInt64([][]int64{
		[]int64{1, 2, 3},
		[]int64{4, 5, 6},
		[]int64{7, 8, 9},
	})
	b := lib.NewMatrixInt64([][]int64{
		[]int64{1, 2, 3},
		[]int64{4, 5, 6},
		[]int64{7, 8, 9},
	})
	fmt.Println(a.Plus(b))
}
