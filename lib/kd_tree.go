package lib

import (
	"fmt"
	"log"
	"math"
	"reflect"

	"mlt/utils"
)

// kd-tree
type KdTree struct {
	Left, Right, Parent *KdTree
	Dims                int // dimensions count
	SplitIndex          int
	SplitValue          float64
	NodeData            Vector
}

//
func NewKdTree(dataSet []Vector) *KdTree {
	if len(dataSet) == 0 {
		return nil
	}
	// check if data set is valid
	if dataSet[0] == nil {
		log.Println("data set must not contain any nil vector")
		return nil
	}
	// find k
	var (
		tree   *KdTree = &KdTree{}
		maxVar float64
		k      int
		dims   int = dataSet[0].Dimensions()
		size   int = len(dataSet)
	)
	for i := 0; i < dims; i++ {
		data := []float64{}
		for j := 0; j < size; j++ {
			data = append(data, dataSet[j].Value(i).(float64))
		}
		_var := utils.VarianceFloat64(data)
		if _var > maxVar {
			maxVar = _var
			k = i
		}
	}
	// find median
	var (
		median    float64 = medianFloat64(dataSet, k) // median value
		medianIdx int                                 // median index
	)
	for i := 0; i < len(dataSet); i++ {
		if dataSet[i].Value(k).(float64) == median {
			medianIdx = i
			break
		}
	}
	// finish tree
	tree.NodeData = dataSet[medianIdx]
	tree.Split = k
	tree.Dims = dims
	var leftDataSet, rightDataSet = []Vector{}, []Vector{} // left data set and right data set
	for i := 0; i < size; i++ {
		if i != medianIdx {
			_v := dataSet[i].Value(k).(float64)
			if _v > median {
				rightDataSet = append(rightDataSet, dataSet[i])
			} else {
				leftDataSet = append(leftDataSet, dataSet[i])
			}
		}
	}
	tree.Left = NewKdTree(leftDataSet)
	tree.Left.Parent = tree
	tree.Right = NewKdTree(rightDataSet)
	tree.Right.Parent = tree
	return tree
}

// find median (type: float64)
func medianFloat64(a []Vector, k int) float64 {
	leng := len(a)
	_a := make([]Vector, leng)
	copy(_a, a)
	return quick(_a, 0, leng-1, leng, k)
}

// quick find
func quick(a []Vector, left, right, leng, k int) float64 {
	p := partition(a, left, right, k)
	if p > leng/2 {
		return quick(a, left, p-1, leng, k)
	} else if p < leng/2 {
		return quick(a, p+1, right, leng, k)
	} else {
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

// search value in kd-tree
func (this *KdTree) Search(v Vector) (*KdTree, float64) {
	// invalid dimension count
	if v.Dimensions() != this.Dims {
		return nil, math.MaxFloat64
	}
	var (
		node  *KdTree = this
		stack Stack   = NewStackKdTree(node)
	)
	// search v until reach leaf node
	for {
		var vSplit float64 = v.Value(node.SplitIndex).(float64)
		stack.Push(node)
		if node.SplitValue > vSplit {
			if node.Left == nil {
				break
			}
			node = node.Left
		} else {
			if node.Right == nil {
				break
			}
			node = node.Right
		}
	}
	// record
	var (
		nearest  *KdTree = node
		distance float64 = utils.L2Norm(v, node.NodeData, reflect.Float64.String())
	)
	// find backward
	stack.Pop()
	for !stack.IsEmpty() {
		_node := node
		node = stack.Pop()

		if _dist := utils.L2Norm(v, node.NodeData, reflect.Float64.String()); _dist < distance {

			// update
			nearest = node
			distance = _dist

			// is parent
			if _node.Parent == node {

				// cross
				if distance > math.Abs(v.Value(node.SplitIndex).(float64)-node.SplitValue) {
					//
					if _node == node.Left {
						_node = node.Right
					} else {
						_node = node.Left
					}
					// now _node is another child of node
					_nnode := &KdTree{
						Left:       _node.Left,
						Right:      _node.Right,
						SplitIndex: _node.SplitIndex,
						SplitValue: _node.SplitValue,
					}
					subNearest, subDistance := _node_node.Search(v)
					if subDistance < distance {
						nearest, distance = subNearest, subDistance
					}
					if stack.IsEmpty() {
						return nearest, distance
					}
					stack.Pop()
				}
			}
		}
	}
	return nearest, distance
}
