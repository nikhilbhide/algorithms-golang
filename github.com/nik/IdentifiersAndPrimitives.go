package main

import "fmt"

//Declare a Variable gloabalvar of type int
var globalVar int
var globalVarString string

func main() {
	//Print the default value of global variables - globalVar & globalVarString
	fmt.Println("Default Value of globalVar before assigning any explicit value :", globalVar)
	fmt.Println("Default Value of globalVarString before assigning any explicit value :", globalVarString, " end")

	x := 10
	fmt.Println("Value of x before :", x)
	x = 20
	fmt.Println("Value of x after :", x)

	var y = "This is my String"
	fmt.Println("Value of y :", y)
	z := 60.20

	//Print the data types of variables - x, y & z
	fmt.Printf("%T\n", z)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", x)
	globalVar = 100
	fmt.Println("Value of globalVar : ", globalVar)

	//Invoke method giveMeX from main
	giveMeX()

	//Assign value to globalVarString variable, which is of type String
	globalVarString = "GoLang Tutorial"
	fmt.Println("Default Value of globalVarString after assigning any explicit value :", globalVarString)

}

func giveMeX() {
	var x = "Golang"
	fmt.Println("Value of x : ", x)

	//Change the value of globalVar to 200
	globalVar = 200
	fmt.Println("Value of globalVar : ", globalVar)
}
