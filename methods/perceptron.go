package methods

import (
	"fmt"
	"log"

	"mlt/lib"
)

// y = sign(w*x+b)
type BasePerceptron struct {
	w lib.Vector
	b float64

	eta float64
}

func signal(input interface{}) int {
	switch input.(type) {
	case float64:
		v := input.(float64)
		if v > 0 {
			return 1
		} else if v < 0 {
			return -1
		}
	case int64:
		v := input.(int64)
		if v > 0 {
			return 1
		} else if v < 0 {
			return -1
		}
	}
	return 0
}

// training  (treat as float64)
func (this *BasePerceptron) Train(xSet []lib.Vector, ySet []int) {
	if len(xSet) == 0 || len(xSet) != len(ySet) {
		return
	}
	for i := 0; i < len(xSet); i++ {
		y, err := this.Predict(xSet[i])
		if err != nil {
			log.Println(err.Error())
			continue
		}
		for y != ySet[i] {
			log.Printf("w=%v b=%10f x=%v\n", this.w, this.b, xSet[i])
			// wrong, fix w and b
			tmpW := this.w.Plus(xSet[i].MultiNumber(this.eta * float64(ySet[i])))
			if tmpW == nil {
				log.Printf("w: %v x: %v eta: %5f y: %d\n", this.w, xSet[i], this.eta, ySet[i])
				continue
			} else {
				this.w = tmpW
				this.b = this.b + this.eta*float64(ySet[i])
			}
			y, err = this.Predict(xSet[i])
			if err != nil {
				log.Println(err.Error())
				break
			}
		}
	}
}

func (this *BasePerceptron) Predict(x lib.Vector) (y int, err error) {
	tmp := this.w.Multi(x)
	if tmp == nil {
		err = fmt.Errorf("invalid dimensions of %v\n", x.Values())
	} else {
		y = signal(tmp.(float64) + this.b)
	}
	return
}

func NewBasePerceptron(w lib.Vector, b, eta float64) *BasePerceptron {
	return &BasePerceptron{w: w, b: b, eta: eta}
}
