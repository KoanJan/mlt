package methods

import (
	"reflect"

	"mlt/lib"
)

// HMM - Hidden Markov Model
type HMM struct {
	pi     lib.Vector
	_A, _B lib.Matrix
}

// Calc and get the probability of the sequence
func (this *HMM) Calc(o []int) float64 {

	// Forward algorithm
	var (
		t    int = len(o)
		_, n int = this._A.Size()
		m, _ int = this._B.Size()

		alpha [][]float64 = make([][]float64, t)
		p     float64
	)
	if t < 1 || t >= m {
		return 0
	}
	for i := 0; i < m; i++ {
		alpha[i] = make([]float64, n)
	}
	// init
	for i := 0; i < n; i++ {
		alpha[0][i] = this.pi.Value(i).(float64) * this._B.Get(i, o[0]).(float64)
	}
	// iter
	for _t := 0; _t < t-1; _t++ {
		for i := 0; i < m; i++ {
			var alpha_t float64
			for j := 0; j < n; j++ {
				alpha_t += alpha[_t][j] * this._A.Get(j, i).(float64)
			}
			alpha[_t+1][i] = alpha_t * this._B.Get(i, o[_t+1]).(float64)
		}
	}
	// end
	for i := 0; i < n; i++ {
		p += alpha[t-1][i]
	}
	return p

	// Backward algorithm
}

// TODO: Train
func (this *HMM) Train() {

	// Baum-Welch algorithm
}

// NewHMM returns an HMM
func NewHMM(pi lib.Vector, a, b lib.Matrix) *HMM {
	if pi == nil || a == nil || b == nil {
		panic("each parameter cannot be nil")
	}
	if pi.Type() != reflect.Float64.String() {
		panic("invalid type of pi")
	}
	if a.Type() != reflect.Float64.String() {
		panic("invalid type of a")
	}
	if b.Type() != reflect.Float64.String() {
		panic("invalid type of b")
	}
	_, n := a.Size()
	m, _ := b.Size()
	if m != n {
		panic("the number of columns of a is not equal with the one of rows of b")
	}
	if pi.Dimensions() != n {
		panic("the number of dimensions of pi is invalid")
	}
	return &HMM{pi: pi, _A: a, _B: b}
}
