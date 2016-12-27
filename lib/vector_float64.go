package lib

import (
	"reflect"
)

type VectorFloat64 struct {
	v []float64
}

func (this *VectorFloat64) Dimensions() int {
	return len(this.v)
}

func (this *VectorFloat64) Type() string {
	return reflect.Float64.String()
}

func (this *VectorFloat64) Value(index int) interface{} {
	if index < 0 || index > len(this.v)-1 {
		return nil
	}
	return this.v[index]
}

func (this *VectorFloat64) Values() []interface{} {
	vals := make([]interface{}, len(this.v))
	for i := 0; i < len(this.v); i++ {
		vals[i] = this.v[i]
	}
	return vals
}

func (this *VectorFloat64) Plus(v Vector) Vector {
	if v.Dimensions() != len(this.v) {
		return nil
	}
	switch v.Type() {
	case "float64":
		var (
			data []float64     = make([]float64, len(this.v))
			vd   []interface{} = v.Values()
		)
		copy(data, this.v)
		for i := 0; i < len(data); i++ {
			data[i] += vd[i].(float64)
		}
		return &VectorFloat64{v: data}
	case "int64":
		var (
			data []float64     = make([]float64, len(this.v))
			vd   []interface{} = v.Values()
		)
		copy(data, this.v)
		for i := 0; i < len(data); i++ {
			data[i] += float64(vd[i].(int64))
		}
		return &VectorFloat64{v: data}
	default:
		return nil
	}
}

func (this *VectorFloat64) Multi(v Vector) interface{} {
	if len(this.v) != v.Dimensions() {
		return nil
	}

	if v.Type() == "int64" {
		var res float64
		for i := 0; i < len(this.v); i++ {
			res += this.v[i] * float64(v.Value(i).(int64))
		}
		return res
	} else {
		var res float64
		for i := 0; i < len(this.v); i++ {
			res += this.v[i] * v.Value(i).(float64)
		}
		return res
	}
}

func (this *VectorFloat64) MultiNumber(n float64) Vector {
	var data []float64 = make([]float64, len(this.v))
	copy(data, this.v)
	for i := 0; i < len(data); i++ {
		data[i] *= n
	}
	return &VectorFloat64{v: data}
}

func NewVectorFloat64(v []float64) Vector {
	return &VectorFloat64{v: v}
}
