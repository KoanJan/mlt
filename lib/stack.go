package lib

type Stack interface {
	Type() string
	Push(interface{})
	Pop() interface{}
	IsEmpty() bool
}
