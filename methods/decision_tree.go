package methods

import (
	"math"

	"mlt/lib"
)

// DT generating algorithm
const (
	DT_GA_ID3  int = 1
	DT_GA_C4_5 int = 2
	DT_GA_CART int = 3
)

// DT postPruning programs
const (
	DT_PRUNING_NO  int = 0
	DT_PRUNING_REP int = 1
	DT_PRUNING_PEP int = 2
	DT_PRUNING_CCP int = 3
)

// DT decision tree
type DT struct {
	ga          int
	postPruning int

	alpha float64

	tree *decisionTree
}

// decisionTree
// If the tree is node, its children or feature shouldn't be nil and values should be nil.
// And on the contrary, if the tree is leaf, its children and feature should be nil and values shouldn't be nil
type decisionTree struct {
	// node
	children map[interface{}]*decisionTree
	feature  *struct {
		index int
		_type string
	}
	// leaf
	values [][]*lib.Value
}

// TODO
func (this *DT) Train(data [][]*lib.Value) {
	//
}

// gaId3 generates a decisionTree with ID3 algorithm
func gaID3(tree *decisionTree, alpha float64, data [][]*lib.Value) {
	// amount of records in data
	m := len(data)
	if m == 0 {
		return
	}
	// amount of features in each record, and last dimension of record is y
	n := len(data[0])
	if n == 0 {
		return
	}
	var (
		ent float64   = entropy((p(data, n-1))...)
		gs  []float64 = make([]float64, n-1)
	)
	// calculate each infomation gain
	for i := 0; i < n-1; i++ {
		var (
			fx map[interface{}]int                 = make(map[interface{}]int)
			fy map[interface{}]map[interface{}]int = make(map[interface{}]map[interface{}]int)
		)
		for j := 0; j < m; j++ {
			fx[data[j][i].V()] += 1
			if _, existed := fy[data[j][i].V()]; !existed {
				fy[data[j][i].V()] = map[interface{}]int{}
			}
			fy[data[j][i].V()][data[n-1][i].V()] += 1
		}
		// conditional entropy
		var h float64
		for kx, vx := range fx {
			proby := []float64{}
			for _, vy := range fy[kx] {
				proby = append(proby, float64(vy)/float64(vx))
			}
			h += ent - (float64(vx)/float64(m))*entropy(proby...)
		}
		gs[i] = ent - h
	}
	// select a feature with max infomation gain
	var (
		featureIndex int     = 0
		maxGain      float64 = gs[0]
	)
	for i := 1; i < n-1; i++ {
		if gs[i] > maxGain {
			maxGain = gs[i]
			featureIndex = i
		}
	}
	// classify
	// TODO
}

func p(data [][]*lib.Value, index int) []float64 {
	var (
		leng int = len(data)
		m    map[interface{}]int
		prob []float64 = []float64{}
	)
	for i := 0; i < leng; i++ {
		m[data[i][index].V()] += 1
	}
	for _, v := range m {
		p = append(p, float64(v)/float64(leng))
	}
	return p
}

// entropy calculates the infomation entropy
func entropy(p ...float64) float64 {
	var e float64
	for i := 0; i < len(p); i++ {
		// set 0log0 with 0
		if p[i] > 0 {
			e -= p[i] * math.Log2(p[i])
		}
	}
	return e
}

// TODO
func (this *DT) Predict(v []*lib.Value) (interface{}, string) {
	//
	return nil, ""
}

// NewDT returns a new DT
func NewDT(ga, postPruning int, alpha float64) *DT {
	return &DT{ga: ga, postPruning: postPruning, alpha: alpha, tree: new(decisionTree)}
}
