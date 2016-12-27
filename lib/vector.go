package lib

type Vector interface {
	Dimensions() int       // count Dimensions
	Type() string          // elem type
	Value(int) interface{} // get value by index
	Values() []interface{} // get all values formated as array

	// calc
	Plus(Vector) Vector         // vector1 + vector2
	Multi(Vector) interface{}   // vector1 * vector2
	MultiNumber(float64) Vector // vector * number
}
