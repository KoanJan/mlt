package methods

import (
	"mlt/lib"
)

type KNN struct {
	tree *lib.KdTree
}

func NewKNN(data []lib.Vector) *KNN {
	return &KNN{tree: lib.NewKdTree(data)}
}

// TODO
func (this *KNN) Search(v lib.Vector, k int) []lib.Vector {

}
