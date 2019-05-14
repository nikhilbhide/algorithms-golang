package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wgInst sync.WaitGroup
var counter int64

func main() {

	for outerloop := 0; outerloop < 1000; outerloop++ {
		counter = 0
		wgInst.Add(20)
		for i := 0; i < 20; i++ {
			go func() () {
				for j := 0; j < 20; j++ {
					atomic.AddInt64(&counter,1)
					fmt.Println(atomic.LoadInt64(&counter))
				}
				wgInst.Done()
			}()
		}
		wgInst.Wait()
		fmt.Println("The final value of counter ", counter)
	}
	//fmt.Println("The number of go routines after wait is ", runtime.NumGoroutine())

}
