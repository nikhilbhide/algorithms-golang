package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wgInstance sync.WaitGroup

func computeComplexOperation1 () {
	for counter:=0 ; counter< 100 ; counter ++ {
		fmt.Println("This is the loop with counter - 1 and current iteration is :", counter)
	}
	wgInstance.Done()
}

func computeComplexOperation2 () {
	for counter:=0 ; counter< 100 ; counter ++ {
		fmt.Println("This is the loop with counter - 2 and current iteration is :", counter)
	}
}

func main() {
	wgInstance.Add(2)
	go computeComplexOperation1() //thread-2 is handling this operation
	go computeComplexOperation2() //thread-3 is handling this operation
	fmt.Println("Before calling a wait : ",runtime.NumGoroutine())
	wgInstance.Wait()
	fmt.Println("After calling a wait : ",runtime.NumGoroutine())
	fmt.Println(runtime.NumGoroutine())
}
