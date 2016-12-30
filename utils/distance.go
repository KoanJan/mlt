package utils

import (
	"fmt"
	"math"
	"reflect"
)

func L2Norm(a, b []interface{}, _type string) float64 {
	if len(a) != len(b) {
		panic("invalid params a and b in mlt.utils.L2Norm")
	}
	var (
		v    float64
		leng int = len(a)
	)
	switch _type {
	case reflect.Float64.String():
		for i := 0; i < leng; i++ {
			v += (a[i].(float64) - b[i].(float64)) * *2
		}
	case reflect.Int64.String():
		for i := 0; i < leng; i++ {
			v += (float64(a[i].(int64)) - float64(b[i].(int64))) * *2
		}
	default:
		panic(fmt.Sprintf("unknown element type of params '%s' in mlt.utils.L2Norm", _type))
	}
	return math.Sqrt(v)
}
