package main

import "fmt"

func main() {
	sample_value:=-9999
	invoke_func_with_int_pointer(100,hello_world_simple)
	func(x *int) {
		fmt.Println("Passed memory location is : ",x)
		fmt.Println("Value associated with the memory address is this : ",*x)
	} (&sample_value)

	sample_function:=func(x *int) {
		fmt.Println("Passed memory location is : ",x)
		fmt.Println("Value associated with the memory address is this : ",*x)
	}

	fmt.Printf("%T\n",sample_function)
	fmt.Printf("%T\n",&sample_function)
	fmt.Println("Memory address of the function - sample_function ",&sample_function)
	invoke_func_with_int_pointer(sample_value,sample_function) //sample_function(sample_value)
}

func hello_world_simple (value *int) {
	fmt.Println("Hello World")
	fmt.Println("Passed value to me is : ",value)
	*value = 10000
}

func invoke_func_with_int_pointer (value int, func_instance func (*int)) {
	fmt.Println("Value passed to me is : ", value)
	fmt.Println("I am invoking a function")
	func_instance(&value)
	fmt.Println("New value is : ", value)
}
