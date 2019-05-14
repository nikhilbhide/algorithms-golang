package main

import (
	"algorithms-golang/github.com/nik/Sample/HelloWorld"
	"algorithms-golang/github.com/nik/Sample/HelloWorld/App"
	"fmt"
)
func tryMe() (int, func (int,int)) {
	counter:=100
	fmt.Println(counter)
	fmt.Println("Counter inside try me : ",&counter)

	funcVariable:= func (a int, b int) {
		fmt.Println("Sum is : ",a+b)
		fmt.Println("counter before is : ",counter)

		counter = 200
		fmt.Println("counter after is : ",counter)
	}

	func (a int, b int) {
		fmt.Println("Sum is : ",a+b)
	} (10,20)

	funcVariable(10,50)
	fmt.Println("counter after invocation is : ",counter)

	return counter,funcVariable
}

func main() {
	fmt.Println(HelloWorld.AppName)
	fmt.Println(HelloWorld.GetAppName())
	fmt.Println(HelloWorld.GetDisplayMessage())
	App.GetSampleApp()

	//counter value
	counter, funcVar:= tryMe()
	fmt.Println("Counter inside main: ",&counter)
	fmt.Println(counter)
	fmt.Println("Peak of my scoping session")
	funcVar(-10,30)
}
