package utils

func VarianceFloat64(data []float64) float64 {
	if len(data) < 2 {
		return 0
	}
	var (
		aver  = AverageFloat64(data)
		_var  float64
		count int = len(data)
	)
	for i := 0; i < count; i++ {
		_var += (aver - data[i]) * (aver - data[i])
	}
	return _var / count
}

func AverageFloat64(data []float64) float64 {
	var (
		sum   float64
		count int = len(data)
	)
	for i := 0; i < count; i++ {
		sum += data[i]
	}
	return sum / count
}
