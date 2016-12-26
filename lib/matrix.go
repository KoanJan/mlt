package lib

type Matrix interface {
	Size() (int, int)          // return rows count and columns count of the matrix
	Type() string              // type of elem
	Get(int, int) interface{}  // get element value
	Set(int, int, interface{}) // set element value
	String() string            // format output as string
	IsEqual(Matrix) bool       // is equal to another matrix

	// calc
	Plus(Matrix) Matrix  // matrix plus
	Multi(Matrix) Matrix // matrix multi
}

func canMatrixPlus(a, b Matrix) bool {
	if a == nil || b == nil {
		return false
	}
	// can translate
	if a.Type() != b.Type() {
		return false
	}
	var (
		r1, c1 int = a.Size()
		r2, c2 int = b.Size()
	)
	return r1 == r2 && c1 == c2
}

func canMatrixMulti(a, b Matrix) bool {
	if a == nil || b == nil {
		return false
	}
	// can translate
	if a.Type() != b.Type() {
		return false
	}
	var (
		_, col1 int = a.Size()
		row2, _ int = b.Size()
	)
	return col1 == row2
}
