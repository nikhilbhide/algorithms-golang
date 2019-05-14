package main

import "fmt"

func main() {
	sumWithLooping(11)
	result_sum := sumWithRecursion(11,0)
	fmt.Println(result_sum)
}

//calculate the sum from 0 to n - recursive function
func sumWithRecursion(n int,current_num int) int {
	if(current_num< n) {
		return current_num + (sumWithRecursion(n,current_num+1))
		//unfold the recursive invocation results - 0 + (1 + (2 + (3 + (4 + (5 + (6 + (7 + (8 + (9 + (10)))))))
	} else {
		return 0
	}
}

//calculate the sum from 0 to n - nonrecursive style
func sumWithLooping(n int) int {
	sum:= 0
	for counter:=0; counter < n; counter++ {
		sum+=counter
	}

	fmt.Println("The sum using looping is : ",sum)
	return sum
}
