package lib

type StackVector struct {
	data []Vector
}

func (this *StackVector) Type() string {
	return "mlt.lib.Vector"
}

func (this *StackVector) Push(v interface{}) {
	if vv, yes := v.(Vector); yes {
		this.data = append(this.data, vv)
	}
}

func (this *StackVector) Pop() interface{} {
	l := len(this.data)
	v := this.data[l-1]
	this.data = this.data[0 : l-1]
	return v
}

func NewStackVector(data ...Vector) Stack {
	stack := &StackVector{data: []Vector{}}
	if len(data) > 0 {
		stack.data = data
	}
	return stack
}
