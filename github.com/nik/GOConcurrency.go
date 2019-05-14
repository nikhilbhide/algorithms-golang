package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())


}
