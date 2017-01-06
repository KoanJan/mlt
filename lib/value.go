package lib

type Value struct {
	v interface{}
	t string
}

func (this *Value) V() interface{} {
	return this.v
}

func (this *Value) T() string {
	return this.t
}

func NewValue(v interface{}, _type string) *Value {
	return &Value{v: v, t: _type}
}
