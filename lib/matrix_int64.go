package lib

import (
	"fmt"
	"reflect"
)

type MatrixInt64 struct {
	data [][]int64
}

func (this *MatrixInt64) Type() string {
	return reflect.Int64.String()
}

func (this *MatrixInt64) Size() (int, int) {
	row := len(this.data)
	if row > 0 {
		return row, len(this.data[0])
	} else {
		return 0, 0
	}
}

func (this *MatrixInt64) Get(row, col int) interface{} {
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
func (this *MatrixInt64) get(row, col int) int64 {
	return this.data[row][col]
}

func (this *MatrixInt64) Set(row, col int, v interface{}) {
	if r, c := this.Size(); row < r && row >= 0 && col < c && col >= 0 {
		if val, yes := v.(int64); yes {
			this.data[r][c] = val
		}
	}
}

func (this *MatrixInt64) String() string {
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
				tmp = fmt.Sprintf("%s%10d ", tmp, this.get(i, j))
			}
			out += tmp + "\n"
		}
	}
	return out
}

func (this *MatrixInt64) IsEqual(m Matrix) bool {
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
			if this.get(i, j) != m.Get(i, j).(int64) {
				return false
			}
		}
	}
	return true
}

func (this *MatrixInt64) Plus(m Matrix) Matrix {
	if canMatrixPlus(this, m) {
		row, col := this.Size()
		data := [][]int64{}
		for i := 0; i < row; i++ {
			r := []int64{}
			for j := 0; j < col; j++ {
				r = append(r, this.get(i, j)+m.Get(i, j).(int64))
			}
			data = append(data, r)
		}
		return NewMatrixInt64(data)
	}
	return nil
}

func (this *MatrixInt64) Multi(m Matrix) Matrix {
	if canMatrixMulti(this, m) {
		var (
			row1, col1 int       = this.Size()
			_, col2    int       = m.Size()
			data       [][]int64 = [][]int64{}
		)
		for i := 0; i < row1; i++ {
			newRow := []int64{}
			for j := 0; j < col2; j++ {
				// col1 is equal to row2
				var v int64
				for k := 0; k < col1; k++ {
					v += this.get(i, k) * m.Get(k, j).(int64)
				}
				newRow = append(newRow, v)
			}
			data = append(data, newRow)
		}
		return NewMatrixInt64(data)
	}
	return nil
}

func NewMatrixInt64(dataInt64 [][]int64) Matrix {
	return &MatrixInt64{data: dataInt64}
}
