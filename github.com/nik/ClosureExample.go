package main

import "fmt"

func main() {
	func_var:=closure_sample_func(100)
	fmt.Printf("%T\n",func_var)
	fmt.Printf("%T\n",&func_var)
	fmt.Println("The memory address of the function variable is : ",&func_var)

	func_var(1000) //this is first call
	func_var(1000) //this is second call
	func_var(1000) //this is third call
	func_var(1000) //this is fourth call
	func_var(1000) //this is fifth call
}

func closure_sample_func(value int) (func (int) int) {
	x:=100
	fmt.Println(x)

	return func(value int) (int) {
		fmt.Println("The value of the passed variable : ", value)
		fmt.Println("The value of the x variable : ", x)
		 x=x+1
		return x
	}
}
