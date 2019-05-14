package errorhandling

import (
	"errors"
	"fmt"
	"math"
)

//business rules and business errors -> audit
//system errors and logs
var err = errors.New("Divide by zero error")

func Divide (a int, b int) (float64, error) {
	if b == 0 {
		fmt.Printf("%T\n",err)
		return math.MaxFloat64,fmt.Errorf("Divide by zero error because parameter b has a value %v",b)
	} else {
		return float64(a)/float64(b),nil
	}
}
