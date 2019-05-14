package main

import (
	"fmt"
)

func main() {
	//declare a variable to hold integer value 100
	value := 100
	fmt.Println("Value is : ", value)           //value gives back 100
	fmt.Println("Memory address is : ", &value) //Where in the memory the variable is located? That's the memory address
	fmt.Println("Dereferencing the address that is get value from address : ", *(&value))

	//check the type difference
	fmt.Printf("%T\n", value)
	fmt.Printf("%T\n", &value)
	fmt.Printf("%T\n", *&value)

	//pass the value to the function as a value and check whether the value is changed
	modify_with_value(value)
	//check whether the value is changed
	fmt.Println("After function invocation new value is : ", value)

	//pass the pointer to a function and check whether the value is changed
	modify_with_pointer(&value)
	//check whether the value is changed
	fmt.Println("After function invocation new value is : ", value)
	//Is there any change in the address? -  NO
	fmt.Println("After function invocation memory address is : ", &value)
}

//function that takes integer
func modify_with_value(value int) {
	fmt.Println("Passed value is : ", value)
	value = 100000
	fmt.Println("Memory address is : ", &value) //Where in the memory the variable is located? That's the memory address
	fmt.Println("Changed value is : ", value)
}

//function that takes integer and changes the value
func modify_with_pointer(value *int) {
	fmt.Println("Passed value is : ", value)
	*value = 100000
}