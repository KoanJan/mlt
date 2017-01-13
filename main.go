package main

import (
	"log"

	"mlt/lib"
	"mlt/methods"
)

func main() {
	// type
	const (
		BYTE string = "byte"
		BOOL string = "bool"
	)
	// age
	const (
		AGE_Y byte = 'y' // youth
		AGE_M byte = 'm' // middle aged
		AGE_E byte = 'e' // elder
	)
	// belief
	const (
		BELIEF_PERFECT  byte = 'p' // perfect
		BELIEF_GOOD     byte = 'g' // good
		BELIEF_ORDINARY byte = 'o' // ordinary
		BELIEF_BAD      byte = 'b' // bad
	)
	// alia for short
	nv := lib.NewValue

	// data
	// record: age, hasJob, hasHouse, belief
	data := [][]*lib.Value{
		[]*lib.Value{nv(AGE_Y, BYTE), nv(false, BOOL), nv(false, BOOL), nv(BELIEF_ORDINARY, BYTE), nv(false, BOOL)},
		[]*lib.Value{nv(AGE_Y, BYTE), nv(false, BOOL), nv(false, BOOL), nv(BELIEF_GOOD, BYTE), nv(false, BOOL)},
		[]*lib.Value{nv(AGE_Y, BYTE), nv(true, BOOL), nv(false, BOOL), nv(BELIEF_GOOD, BYTE), nv(true, BOOL)},
		[]*lib.Value{nv(AGE_Y, BYTE), nv(true, BOOL), nv(true, BOOL), nv(BELIEF_ORDINARY, BYTE), nv(true, BOOL)},
		[]*lib.Value{nv(AGE_Y, BYTE), nv(false, BOOL), nv(false, BOOL), nv(BELIEF_ORDINARY, BYTE), nv(false, BOOL)},

		[]*lib.Value{nv(AGE_M, BYTE), nv(false, BOOL), nv(false, BOOL), nv(BELIEF_ORDINARY, BYTE), nv(false, BOOL)},
		[]*lib.Value{nv(AGE_M, BYTE), nv(false, BOOL), nv(false, BOOL), nv(BELIEF_GOOD, BYTE), nv(false, BOOL)},
		[]*lib.Value{nv(AGE_M, BYTE), nv(true, BOOL), nv(true, BOOL), nv(BELIEF_GOOD, BYTE), nv(true, BOOL)},
		[]*lib.Value{nv(AGE_M, BYTE), nv(false, BOOL), nv(true, BOOL), nv(BELIEF_PERFECT, BYTE), nv(true, BOOL)},
		[]*lib.Value{nv(AGE_M, BYTE), nv(false, BOOL), nv(true, BOOL), nv(BELIEF_PERFECT, BYTE), nv(true, BOOL)},

		[]*lib.Value{nv(AGE_E, BYTE), nv(false, BOOL), nv(true, BOOL), nv(BELIEF_PERFECT, BYTE), nv(true, BOOL)},
		[]*lib.Value{nv(AGE_E, BYTE), nv(false, BOOL), nv(true, BOOL), nv(BELIEF_GOOD, BYTE), nv(true, BOOL)},
		[]*lib.Value{nv(AGE_E, BYTE), nv(true, BOOL), nv(false, BOOL), nv(BELIEF_GOOD, BYTE), nv(true, BOOL)},
		[]*lib.Value{nv(AGE_E, BYTE), nv(true, BOOL), nv(false, BOOL), nv(BELIEF_PERFECT, BYTE), nv(true, BOOL)},
		[]*lib.Value{nv(AGE_E, BYTE), nv(false, BOOL), nv(false, BOOL), nv(BELIEF_ORDINARY, BYTE), nv(false, BOOL)},

		[]*lib.Value{nv(AGE_Y, BYTE), nv(false, BOOL), nv(false, BOOL), nv(BELIEF_PERFECT, BYTE), nv(true, BOOL)},
		[]*lib.Value{nv(AGE_Y, BYTE), nv(false, BOOL), nv(true, BOOL), nv(BELIEF_BAD, BYTE), nv(false, BOOL)},
	}

	// create a decision tree
	dt := methods.NewDT(methods.DT_GA_C4_5, methods.DT_PRUNING_NO, 0.5)

	// test a new record
	v := []*lib.Value{nv(AGE_E, BYTE), nv(true, BOOL), nv(false, BOOL), nv(BELIEF_BAD, BYTE)}

	// train model
	dt.Train(data)

	// make predict
	class, classType := dt.Predict(v)

	log.Printf("%v %s", class, classType)
}
