package main

import (
	"fmt"
	"runtime"
)

func main() {
	//display the architecture details
	fmt.Println("Architecture of machine is :", runtime.GOARCH)
	//display os details
	fmt.Println("Os details : ", runtime.GOOS)
	//display the number of processors
	fmt.Println("Number of processors : ",runtime.NumCPU())
	//display the number of routines
	fmt.Println("Number of routines : ",runtime.NumGoroutine())
}
