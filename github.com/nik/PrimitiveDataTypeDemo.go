package main

import "fmt"

func main() {
	//Declare a variable and assign value "true" to it
	boolVar := true
	fmt.Println("The value of boolVar is : ", boolVar)
	//Change value to "false"
	boolVar = false
	fmt.Println("The value of boolVar is : ", boolVar)

	//Declare a variable of type uint32 and int32
	var uintVar uint32
	var intVar int32
	//Assign the extreme left value and right value to the uintVar variable and intVar variable
	uintVar = 0
	intVar= -2147483648
	fmt.Println("The extreme left value of unitVar is : ", uintVar)
	fmt.Println("The extreme left value  of intVar is : ", intVar)
	uintVar = 4294967295
	intVar= 2147483647
	fmt.Println("The extreme right value of unitVar is : ", uintVar)
	fmt.Println("The extreme right value  of intVar is : ", intVar)

	//Declare variables of type float32 and float64
	var float32Var float32 = 100.60033678676788
	var float64Var float64 = 100.60033678676788
	fmt.Println("The Value of float32Var is : ", float32Var)
	fmt.Println("The Value of float64Var is : ", float64Var)

	//Declare a variable of type string and represent the string in bytes
	var stringVar string = "I love GoLang programming"
	fmt.Println("The Value of stringVar: ", stringVar)

	byteVarOfString:= [] byte(stringVar)
	fmt.Println("Byte representation of the string is : ",byteVarOfString)
}
