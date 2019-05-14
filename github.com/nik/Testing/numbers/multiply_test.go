package numbers

import (
	"testing"
)

type Test struct {
	num1   int
	num2   int
	num3   []int
	result int
}

func TestMultiply(t *testing.T) {
	testData := [] Test{Test{num1: 10, num2: 20, num3: [] int{20, 2, 3}, result: 1000}}
	//iterate over composite literal
	actualResult:=0
	for index, value := range testData {
		//actualResult := multiplyNumbers(append(value.num3, value.num1, value.num2)...))
		if actualResult != value.result {
			t.Error("Expected: ", value.result, " actual:", actualResult)
			t.Log("Error occured at test data index : ",index)
		} else {
			t.Log("Test case passed at test data index : ",index)
		}
	}
}

