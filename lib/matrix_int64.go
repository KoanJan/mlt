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
		return this.data[row][col]
	}
	return nil
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
				tmp = fmt.Sprintf("%s%10d ", tmp, this.Get(i, j).(int64))
			}
			out += tmp + "\n"
		}
	}
	return out
}

func (this *MatrixInt64) Plus(m Matrix) Matrix {
	if canMatrixPlus(this, m) {
		row, col := this.Size()
		data := [][]int64{}
		for i := 0; i < row; i++ {
			r := []int64{}
			for j := 0; j < col; j++ {
				r = append(r, this.data[i][j]+m.Get(i, j).(int64))
			}
			data = append(data, r)
		}
		return NewMatrixInt64(data)
	}
	return nil
}

func NewMatrixInt64(dataInt64 [][]int64) Matrix {
	return &MatrixInt64{data: dataInt64}
}
