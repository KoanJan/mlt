package methods

import (
	"mlt/lib"
)

// naive Bayesian classifier
type NaiveBayesianClassifier struct {
	xProbabilities []map[interface{}]map[interface{}]float64
	yProbabilities map[interface{}]float64
	recordTypes    []string
	n              int
}

// while training a new data set, it will clear what the classifier has learn last time
func (this *NaiveBayesianClassifier) Train(data [][]*lib.Value) {
	// len of records
	m := len(data)
	if m == 0 {
		return
	}
	// dimensions of each record
	n := len(data[0])
	if n == 0 {
		return
	}
	this.n = n
	// init xProbabilities
	this.xProbabilities = make([]map[interface{}]map[interface{}]float64, n-2)
	for i := 0; i < m-1; i++ {
		this.xProbabilities[i] = map[interface{}]map[interface{}]float64{}
	}
	// init yProbabilities
	this.yProbabilities = map[interface{}]float64{}
	// init record types
	this.recordTypes = make([]string, n)
	for i := 0; i < n; i++ {
		this.recordTypes[i] = data[0][i].T()
	}
	// calc xProbabilities and yProbabilities and store them
	xFreq := make([]map[interface{}]map[interface{}]int, n-1)
	for i := 0; i < n-1; i++ {
		xFreq[i] = map[interface{}]map[interface{}]int{}
	}
	yFreq := make(map[interface{}]float64)
	for i := 0; i < m; i++ {
		for j := 0; j < n-1; j++ {
			if xFreq[j][data[i][j].V()] == nil {
				xFreq[j][data[i][j].V()] = map[interface{}]float64{}
			}
			xFreq[j][data[i][j].V()][data[i][n-1].V()] += 1
		}
		yFreq[data[i][n-1].V()] += 1

	}
	for i := 0; i < n; i++ {
		// Laplace smoothing, and lmd representing λ is set with 1
		lmd := len(yFreq[i])
		for kx, fx := range xFreq[i] {
			if fx == nil {
				panic("invalid data")
			}
			for ky, fy := range fx {
				if this.xProbabilities[i][kx] == nil {
					this.xProbabilities[i][kx] = map[interface{}]float64{}
				}
				this.xProbabilities[i][kx][ky] = float64(fy+1) / (yFreq[ky] + lmd)
			}
		}
	}
	//
}

// predict
func (this *NaiveBayesianClassifier) Predict(v []*lib.Value) (interface{}, string) {
	if len(v) != n {
		return nil, ""
	}
	var (
		fs    map[interface{}]float64 = make(map[interface{}]float64)
		y     interface{}             = nil
		maxYP float64                 = 0
	)
	for ky, vy := range this.yProbabilities {
		fs[ky] = vy
	}
	for i := 0; i < len(v); i++ {
		// nil means that the dimension of whose index doesn't limit value
		if v[i] != nil {
			for ky, _ := range fs {
				fs[ky] *= this.xProbabilities[v[i].V()][ky]
			}
		}
	}
	// select y with max probability
	for ky, vy := range fs {
		if vy > maxYP {
			y = ky
			maxYP = vy
		}
	}
	return y, this.recordTypes[this.n-1]
}
