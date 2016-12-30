package lib

type StackKdTree struct {
	data []*KdTree
}

func (this *StackKdTree) Type() string {
	return "mlt.lib.kd_tree"
}

func (this *StackKdTree) Push(v interface{}) {
	if kd_tree, yes := v.(*KdTree); yes {
		this.data = append(this.data, kd_tree)
	}
}

func (this *StackKdTree) Pop() interface{} {
	l := len(this.data)
	v := this.data[l-1]
	this.data = this.data[0 : l-1]
	return v
}

func (this *StackKdTree) IsEmpty() bool {
	return len(this.data) == 0
}

func NewStackKdTree(data ...KdTree) Stack {
	stack := &StackKdTree{data: []KdTree{}}
	if len(data) > 0 {
		stack.data = data
	}
	return stack
}
