package lib

type Matrix interface {
	Size() (int, int)         // return rows count and columns count of the matrix
	Type() string             // type of elem
	Get(int, int) interface{} // get element value
	String() string           // format output as string

	// calc
	Plus(Matrix) Matrix // matrix plus
	// TODO
	// Multi(Matrix) Matrix // matrix multi
}

func canMatrixPlus(a, b Matrix) bool {
	if a == nil || b == nil {
		return false
	}
	if a.Type() != b.Type() {
		return false
	}
	var (
		r1, c1 int = a.Size()
		r2, c2 int = b.Size()
	)
	return r1 == r2 && c1 == c2
}
