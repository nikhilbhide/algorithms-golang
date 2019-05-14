package main

import "fmt"

//Declare a constant variable and assign a value to it
const constVar = 10

func main() {
	//Declare a constant and do not use it anywhere in the program
	const constNoUse = "NoUse"

	fmt.Println("The value of constant is : ", constVar)
	fmt.Printf("The data type of the constant is : %T\n",constVar)
}