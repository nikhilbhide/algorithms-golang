package main

import (
	"fmt"
)

func main() {

	//The basic switch example without any parameter
	switch {
	case true:
		fmt.Println("This is the true case")
	case false:
		fmt.Println("This is the false case")
	}

	//Parameterized switch statement example
	x := "Hindi"
	switch (x) {
	case "English":
		fmt.Println("The English language case is getting executed")
	case "Hindi":
		fmt.Println("The Hindi language case is getting executed")
	case "Arabic", "Hebrew":
		fmt.Println("Some language from the middle east case is getting executed")
	}

	//Parameterized switch statement example
	language := "French"
	switch (language) {
	case "English":
		fmt.Println("The English language case is getting executed")
	case "Hindi":
		fmt.Println("The Hindi language case is getting executed")
	case "Arabic", "Hebrew":
		fmt.Println("Some language from the middle east case is getting executed")
	default:
		fmt.Println("Some language from the set {French, Sanskrit, African} case is getting executed")
	}

	//Fallthrough switch statement example
	language = "French"
	switch (language) {
	case "English":
		fmt.Println("The English language case is getting executed")
	case "Hindi":
		fmt.Println("The Hindi language case is getting executed")
	case "French":
		fmt.Println("The French language case is getting executed")
		fallthrough
	case "Arabic", "Hebrew":
		fmt.Println("Some language from the middle east case is getting executed")
	default:
		fmt.Println("Some language from the set {Sanskrit, African} case is getting executed")
	}

	//Invoke a typename function
	typeName(language)
}

/*
Declare a function to get the interface value and use switch statement over the type of the value
 */
func typeName(language interface{}) {
	//Type switch statement example with the help of the helper variable
	switch t := language.(type) { //t is the helper variable
	case nil:
		fmt.Printf("The type is %T\n", t)
	case string:
		fmt.Printf("The type is %T\n", t)
	case int:
		fmt.Printf("The type is %T\n", t)
	default:
		fmt.Println("New type!!")
	}

	//Type switch statement example
	switch language.(type) {
	case nil:
		fmt.Printf("The type is nil \n")
	case string:
		fmt.Printf("The type is string\n")
	case int:
		fmt.Printf("The type is int\n")
	default:
		fmt.Println("New type!!")
	}
}