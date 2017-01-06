package methods

import (
	"mlt/lib"
)

// naive Bayesian classifier
type NaiveBayesianClassifier struct {
	probabilities []map[interface{}]float64
	recordTypes   []string
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
	// init probabilities
	this.probabilities = make([]map[interface{}]float64, m-1)
	for i := 0; i < m-1; i++ {
		this.probabilities[i] = map[interface{}]float64{}
	}
	// init record types
	this.recordTypes = make([]string, n)
	for i := 0; i < n; i++ {
		this.recordTypes[i] = data[0][i].T()
	}
	// calc probabilities and store them
	freq := make([]map[interface{}]int, n)
	for i := 0; i < n; i++ {
		freq[i] = map[interface{}]int{}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			freq[j][data[i][j].V()] += 1
		}
	}
	for i := 0; i < n; i++ {
		for k, f := range freq[i] {
			this.probabilities[i][k] = float64(f) / m
		}
	}
	//
}

// predict
func (this *NaiveBayesianClassifier) Predict(v []*lib.Value) (interface{}, string) {

}
