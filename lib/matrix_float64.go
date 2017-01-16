package lib

import (
	"fmt"
	"reflect"
)

type MatrixFloat64 struct {
	data [][]float64
}

func (this *MatrixFloat64) Type() string {
	return reflect.Float64.String()
}

func (this *MatrixFloat64) Size() (int, int) {
	row := len(this.data)
	if row > 0 {
		return row, len(this.data[0])
	} else {
		return 0, 0
	}
}

func (this *MatrixFloat64) Get(row, col int) interface{} {
	if row < 0 || col < 0 {
		return nil
	}
	if r, c := this.Size(); row < r && col < c {
		return this.get(row, col)
	}
	return nil
}

// unsafe.
// before use it you must make sure that row and col is valid, or you'd better use Get
func (this *MatrixFloat64) get(row, col int) float64 {
	return this.data[row][col]
}

func (this *MatrixFloat64) Set(row, col int, v interface{}) {
	if r, c := this.Size(); row < r && row >= 0 && col < c && col >= 0 {
		if val, yes := v.(float64); yes {
			this.data[r][c] = val
		}
	}
}

func (this *MatrixFloat64) String() string {
	var (
		out  string
		r, c int = this.Size()
	)
	if r == 0 && c == 0 {
		out = "<nil>"
	} else {
		for i := 0; i < r; i++ {
			var tmp string
			for j := 0; j < c; j++ {
				tmp = fmt.Sprintf("%s%10f ", tmp, this.get(i, j))
			}
			out += tmp + "\n"
		}
	}
	return out
}

func (this *MatrixFloat64) IsEqual(m Matrix) bool {
	if m == nil {
		return false
	}
	if this.Type() != m.Type() {
		return false
	}
	var (
		r1, c1 int = this.Size()
		r2, c2 int = m.Size()
	)
	if r1 != r2 || c1 != c2 {
		return false
	}
	for i := 0; i < r1; i++ {
		for j := 0; j < c1; j++ {
			if this.get(i, j) != m.Get(i, j).(float64) {
				return false
			}
		}
	}
	return true
}

func (this *MatrixFloat64) Transpose() Matrix {
	var (
		m, n    int         = this.Size()
		newData [][]float64 = [][]float64{}
	)
	for i := 0; i < n; i++ {
		row := []float64{}
		for j := 0; j < m; j++ {
			row = append(row, this.data[j][i])
		}
		newData = append(newData, row)
	}
	return NewMatrixFloat64(newData)
}

func (this *MatrixFloat64) Plus(m Matrix) Matrix {
	if canMatrixPlus(this, m) {
		row, col := this.Size()
		data := [][]float64{}
		for i := 0; i < row; i++ {
			r := []float64{}
			for j := 0; j < col; j++ {
				r = append(r, this.get(i, j)+m.Get(i, j).(float64))
			}
			data = append(data, r)
		}
		return NewMatrixFloat64(data)
	}
	return nil
}

func (this *MatrixFloat64) Multi(m Matrix) Matrix {
	if canMatrixMulti(this, m) {
		var (
			row1, col1 int         = this.Size()
			_, col2    int         = m.Size()
			data       [][]float64 = [][]float64{}
		)
		for i := 0; i < row1; i++ {
			newRow := []float64{}
			for j := 0; j < col2; j++ {
				// col1 is equal to row2
				var v float64
				for k := 0; k < col1; k++ {
					v += this.get(i, k) * m.Get(k, j).(float64)
				}
				newRow = append(newRow, v)
			}
			data = append(data, newRow)
		}
		return NewMatrixFloat64(data)
	}
	return nil
}

func (this *MatrixFloat64) MultiNumber(v interface{}) Matrix {
	var val float64
	if _v, yes := v.(float64); yes {
		val = _v
	} else if _v2, yes2 := v.(int); yes2 {
		val = float64(_v2)
	} else {
		return nil
	}
	var (
		data     [][]float64 = [][]float64{}
		row, col int         = this.Size()
	)
	copy(data, this.data)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			data[i][j] *= val
		}
	}
	return NewMatrixFloat64(data)
}

func NewMatrixFloat64(dataFloat64 [][]float64) Matrix {
	return &MatrixFloat64{data: dataFloat64}
}
