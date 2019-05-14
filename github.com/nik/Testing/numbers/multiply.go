package numbers

import "fmt"

func MultiplyNumbers(numbers ...int) int {
	result := 1
	for _, value := range numbers {
		//fmt.Println(value)
		result = result * value
	}

	if(result==0) {
		fmt.Println("This line is bound to get ignored")
	}
	return result
}
