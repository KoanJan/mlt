package lib

import (
	"reflect"
)

type VectorInt64 struct {
	v []int64
}

func (this *VectorInt64) Dimensions() int {
	return len(this.v)
}

func (this *VectorInt64) Type() string {
	return reflect.Int64.String()
}

func (this *VectorInt64) Value(index int) interface{} {
	if index < 0 || index > len(this.v)-1 {
		return nil
	}
	return this.v[index]
}

func (this *VectorInt64) Values() []interface{} {
	vals := make([]interface{}, len(this.v))
	for i := 0; i < len(this.v); i++ {
		vals[i] = this.v[i]
	}
	return vals
}

func (this *VectorInt64) Plus(v Vector) Vector {
	if v.Dimensions() != len(this.v) {
		return nil
	}
	switch v.Type() {
	case "float64":
		var (
			data []float64     = make([]float64, len(this.v))
			vd   []interface{} = v.Values()
		)
		for i := 0; i < len(data); i++ {
			data[i] += vd[i].(float64) + float64(this.v[i])
		}
		return &VectorFloat64{v: data}
	case "int64":
		var (
			data []int64       = make([]int64, len(this.v))
			vd   []interface{} = v.Values()
		)
		copy(data, this.v)
		for i := 0; i < len(data); i++ {
			data[i] += vd[i].(int64)
		}
		return &VectorInt64{v: data}
	default:
		return nil
	}
}

func (this *VectorInt64) Multi(v Vector) interface{} {
	if len(this.v) != v.Dimensions() {
		return nil
	}

	if v.Type() == "float64" {
		var res float64
		for i := 0; i < len(this.v); i++ {
			res += float64(this.v[i]) * v.Value(i).(float64)
		}
		return res
	} else {
		var res int64
		for i := 0; i < len(this.v); i++ {
			res += this.v[i] * v.Value(i).(int64)
		}
		return res
	}
}

func (this *VectorInt64) MultiNumber(n float64) Vector {
	data := make([]float64, len(this.v))
	for i := 0; i < len(data); i++ {
		data[i] = n * float64(this.v[i])
	}
	return &VectorFloat64{v: data}
}

func NewVectorInt64(v []int64) Vector {
	return &VectorInt64{v: v}
}
