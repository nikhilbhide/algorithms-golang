package main

import (
	"fmt"
	"runtime"
)

func computeComplexOperation_1 () {
	for counter:=0 ; counter< 100 ; counter ++ {
		fmt.Println("This is the loop with counter - 1 and current iteration is :", counter)
	}
}

func computeComplexOperation_2 () {
	for counter:=0 ; counter< 100 ; counter ++ {
		fmt.Println("This is the loop with counter - 2 and current iteration is :", counter)
	}
}

func main() {
	go computeComplexOperation_1() //thread-2 is handling this operation
	go computeComplexOperation_2() //thread-3 is handling this operation

	fmt.Println(runtime.NumGoroutine())
}
