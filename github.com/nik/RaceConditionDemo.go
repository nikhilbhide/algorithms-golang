package main

import (
	"fmt"
	"sync"
)

var wgIns sync.WaitGroup
var mutexIns sync.Mutex

func main() {

	var counter int
	for outerloop := 0; outerloop < 1000; outerloop++ {
		counter = 0
		wgIns.Add(20)
		for i := 0; i < 20; i++ {
			go func() () {

				for j := 0; j < 20; j++ {
					//critical region start
					mutexIns.Lock()
					counter = counter + 1
					mutexIns.Unlock()
					//critical region end
				}
				wgIns.Done()
			}()
		}
		//fmt.Println("The number of go routines before wait is ", runtime.NumGoroutine())
		wgIns.Wait()
		fmt.Println("The final value of counter ", counter)
	}
	//fmt.Println("The number of go routines after wait is ", runtime.NumGoroutine())

}
