package main

import "fmt"

type Test struct {
	num1   int
	num2   int
	num3   []int
	result int
}

func multiplyNumbers(numbers ...int) int {
	result := 1
	for _, value := range numbers {
		result = result * value
	}
	return result
}

func main() {
	testData := [] Test{Test{num1: 10, num2: 20, num3:[] int {20,2,3}, result:1000}, Test{num1: 30, num2: 40, num3:[] int{-10,20}, result: 30,}}

	//iterate over composite literal
	for _, value := range testData {
		fmt.Println(multiplyNumbers(append(value.num3,value.num1, value.num2)...))
	}
}
