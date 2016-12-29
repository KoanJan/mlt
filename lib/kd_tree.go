package lib

import (
	"fmt"
	"log"
)

type KdTree struct {
	Left, Right, Parent *KdTree
	Range               *KdTreeRange
	NodeData            Vector
}

type KdTreeRange struct {
	Index    int
	Max, Min float64
}

func NewKdTree(dataSet []Vector, depth ...int) *KdTree {
	tree := &KdTree{}
	// if len(dataSet) == 0 {
	// 	return tree
	// }
	// // get the depth of node
	// var d int
	// if len(depth) > 0 {
	// 	d = depth[0]
	// }
	// // check if data set is valid
	// if dataSet[0] == nil {
	// 	log.Println("data set must not contain any nil vector")
	// 	return tree
	// }
	// // find split k
	// k := d % dataSet[0].Dimensions()
	// TODO find median
	return tree
}

// find median (type: float64)
func MedianFloat64(a []Vector, k int) float64 {
	log.Printf(stringVectors(a, k))

	leng := len(a)
	_a := make([]Vector, leng)
	copy(_a, a)
	return quick(_a, 1, leng-1, leng, k)
}

// quick find
func quick(a []Vector, left, right, leng, k int) float64 {
	// log.Printf("%d %d\n", left, right)
	p := partition(a, left, right, k)
	// log.Printf("%d\n", p)
	if p > leng/2 {
		return quick(a, left, p-1, leng, k)
	} else if p < leng/2 {
		return quick(a, p+1, right, leng, k)
	} else {
		log.Printf("a: %s l: %d r: %d p: %d\n", stringVectors(a, k), left, right, p)
		return a[p].Value(k).(float64)
	}
}

func partition(a []Vector, left, right, k int) int {
	i := left - 1
	v := a[right].Value(k).(float64)
	for j := left; j < right; j++ {
		if a[j].Value(k).(float64) < v {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[right] = a[right], a[i+1]
	return i + 1
}

func stringVectors(a []Vector, k int) string {
	s := ""
	for i := 0; i < len(a); i++ {
		s += fmt.Sprintf("%v ", a[i].Value(k))
	}
	s += "\n"
	return s
}
