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
	ga          int // generating algorithm
	postPruning int // post pruning program

	threshold float64 // threshold

	classType string // type of class

	tree *decisionTree // decision tree structure
}

// decisionTree
// If the tree is node, its children or feature shouldn't be nil and class should be nil.
// And on the contrary, if the tree is leaf, its children and feature should be nil and class shouldn't be nil
type decisionTree struct {

	// node
	children map[interface{}]*decisionTree
	feature  *decisionTreeFeature

	// leaf
	class interface{}

	isLeaf bool
}

type decisionTreeFeature struct {
	index int
	_type string
}

// train
func (this *DT) Train(data [][]*lib.Value) {
	// init
	this.tree = new(decisionTree)

	// generation
	switch this.ga {
	case DT_GA_ID3:
		gaID3(this.tree, this.threshold, data)
	case DT_GA_C4_5:
		gaC4_5(this.tree, this.threshold, data)
	case DT_GA_CART:
		// TODO
	default:
		panic("unknown generating algorithm")
	}

	// pruning
	switch this.postPruning {
	case DT_PRUNING_NO:
		break
	case DT_PRUNING_REP:
		// TODO
	case DT_PRUNING_PEP:
		// TODO
	case DT_PRUNING_CCP:
		// TODO
	default:
		panic("unknown post pruning program")
	}
}

// gaId3 generates a decisionTree with ID3 algorithm
func gaID3(tree *decisionTree, threshold float64, data [][]*lib.Value) {
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
			fy[data[j][i].V()][data[j][n-1].V()] += 1
		}
		// conditional entropy
		var h float64
		for kx, vx := range fx {
			proby := []float64{}
			for _, vy := range fy[kx] {
				proby = append(proby, float64(vy)/float64(vx))
			}
			h += (float64(vx) / float64(m)) * entropy(proby...)
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
	// end generation
	if maxGain < threshold {
		var (
			maxClass interface{}
			maxCount int
			cm       map[interface{}]int = make(map[interface{}]int)
		)
		for i := 0; i < m; i++ {
			cm[data[i][n-1].V()] += 1
		}
		for k, v := range cm {
			if v > maxCount {
				maxClass = k
				maxCount = v
			}
		}
		tree.class = maxClass
		tree.isLeaf = true
		return
	}
	// select the feature having max classes
	tree.feature = &decisionTreeFeature{index: featureIndex, _type: data[0][featureIndex].T()}
	tree.children = map[interface{}]*decisionTree{}
	// classify
	var subData map[interface{}][][]*lib.Value = make(map[interface{}][][]*lib.Value)
	for i := 0; i < m; i++ {
		feature := data[i][featureIndex].V()
		if subData[feature] == nil {
			subData[feature] = [][]*lib.Value{}
		}
		subData[feature] = append(subData[feature], data[i])
	}
	// recursively call gaID3
	for k, v := range subData {
		tree.children[k] = new(decisionTree)
		gaID3(tree.children[k], threshold, v)
	}
}

// gaC4_5 generates a decisionTree with C4.5 algorithm
func gaC4_5(tree *decisionTree, threshold float64, data [][]*lib.Value) {
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
		grs []float64 = make([]float64, n-1)
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
			fy[data[j][i].V()][data[j][n-1].V()] += 1
		}
		// conditional entropy
		var h float64
		for kx, vx := range fx {
			proby := []float64{}
			for _, vy := range fy[kx] {
				proby = append(proby, float64(vy)/float64(vx))
			}
			h += (float64(vx) / float64(m)) * entropy(proby...)
		}
		grs[i] = (ent - h) / ent
	}
	// select a feature with max infomation gain
	var (
		featureIndex int     = 0
		maxGainR     float64 = grs[0]
	)
	for i := 1; i < n-1; i++ {
		if grs[i] > maxGainR {
			maxGainR = grs[i]
			featureIndex = i
		}
	}
	// end generation
	if maxGainR < threshold {
		var (
			maxClass interface{}
			maxCount int
			cm       map[interface{}]int = make(map[interface{}]int)
		)
		for i := 0; i < m; i++ {
			cm[data[i][n-1].V()] += 1
		}
		for k, v := range cm {
			if v > maxCount {
				maxClass = k
				maxCount = v
			}
		}
		tree.class = maxClass
		tree.isLeaf = true
		return
	}
	// select the feature having max classes
	tree.feature = &decisionTreeFeature{index: featureIndex, _type: data[0][featureIndex].T()}
	tree.children = map[interface{}]*decisionTree{}
	// classify
	var subData map[interface{}][][]*lib.Value = make(map[interface{}][][]*lib.Value)
	for i := 0; i < m; i++ {
		feature := data[i][featureIndex].V()
		if subData[feature] == nil {
			subData[feature] = [][]*lib.Value{}
		}
		subData[feature] = append(subData[feature], data[i])
	}
	// recursively call gaID3
	for k, v := range subData {
		tree.children[k] = new(decisionTree)
		gaC4_5(tree.children[k], threshold, v)
	}
}

// calculate probabilities
func p(data [][]*lib.Value, index int) []float64 {
	var (
		leng int                 = len(data)
		m    map[interface{}]int = make(map[interface{}]int)
		prob []float64           = []float64{}
	)
	for i := 0; i < leng; i++ {
		m[data[i][index].V()] += 1
	}
	for _, v := range m {
		prob = append(prob, float64(v)/float64(leng))
	}
	return prob
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

// classify
func (this *DT) Predict(v []*lib.Value) (interface{}, string) {
	tree := this.tree
	for !tree.isLeaf {
		tree = tree.children[v[tree.feature.index].V()]
	}
	return tree.class, this.classType
}

// NewDT returns a new DT
func NewDT(ga, postPruning int, threshold float64) *DT {
	if threshold == 0 {
		panic("threshold should be larger than 0")
	}
	return &DT{ga: ga, postPruning: postPruning, threshold: threshold, tree: new(decisionTree)}
}
