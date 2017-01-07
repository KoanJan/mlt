package main

import (
	"log"

	"mlt/lib"
	"mlt/methods"
)

func main() {
	// data
	const (
		INT  string = "int"
		BYTE string = "byte"
	)
	data := [][]*lib.Value{
		[]*lib.Value{lib.NewValue(1, INT), lib.NewValue('A', BYTE), lib.NewValue(1, INT)},
		// []*lib.Value{lib.NewValue(2, INT), lib.NewValue('A', BYTE), lib.NewValue(1, INT)},
		[]*lib.Value{lib.NewValue(3, INT), lib.NewValue('A', BYTE), lib.NewValue(1, INT)},
		[]*lib.Value{lib.NewValue(4, INT), lib.NewValue('A', BYTE), lib.NewValue(1, INT)},
		[]*lib.Value{lib.NewValue(5, INT), lib.NewValue('A', BYTE), lib.NewValue(1, INT)},
		[]*lib.Value{lib.NewValue(1, INT), lib.NewValue('B', BYTE), lib.NewValue(1, INT)},
		[]*lib.Value{lib.NewValue(2, INT), lib.NewValue('B', BYTE), lib.NewValue(1, INT)},
		// []*lib.Value{lib.NewValue(3, INT), lib.NewValue('B', BYTE), lib.NewValue(1, INT)},
		[]*lib.Value{lib.NewValue(4, INT), lib.NewValue('B', BYTE), lib.NewValue(1, INT)},
		[]*lib.Value{lib.NewValue(5, INT), lib.NewValue('B', BYTE), lib.NewValue(-1, INT)},
		// []*lib.Value{lib.NewValue(1, INT), lib.NewValue('C', BYTE), lib.NewValue(1, INT)},
		[]*lib.Value{lib.NewValue(2, INT), lib.NewValue('C', BYTE), lib.NewValue(1, INT)},
		[]*lib.Value{lib.NewValue(3, INT), lib.NewValue('C', BYTE), lib.NewValue(1, INT)},
		[]*lib.Value{lib.NewValue(4, INT), lib.NewValue('C', BYTE), lib.NewValue(-1, INT)},
		[]*lib.Value{lib.NewValue(5, INT), lib.NewValue('C', BYTE), lib.NewValue(-1, INT)},
		[]*lib.Value{lib.NewValue(1, INT), lib.NewValue('D', BYTE), lib.NewValue(1, INT)},
		// []*lib.Value{lib.NewValue(2, INT), lib.NewValue('D', BYTE), lib.NewValue(1, INT)},
		[]*lib.Value{lib.NewValue(3, INT), lib.NewValue('D', BYTE), lib.NewValue(-1, INT)},
	}

	// train
	nbClassifier := methods.NewNaiveBayesianClassifier()
	nbClassifier.Train(data)

	// test
	v := []*lib.Value{lib.NewValue(1, INT), lib.NewValue('C', BYTE)}
	y, t := nbClassifier.Predict(v)
	log.Printf("\ninput: {%v %v}\noutput: {value: %v type: %s}\n", v[0].V(), v[1].V(), y, t)
}
