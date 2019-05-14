package main

import "fmt"

func main() {
	//Execute the if block in any case
	if true {
		fmt.Println("This is true")
	}

	//Declare a variable and assign a particular int value to the same
	x := 10
	if x == 10 {
		//Execute only if x has a value of 10
		fmt.Println("The value of x is 10 and condition has been satisfied")
	}

	//Use negation operator in the if
	if false {
		//Execute only if x has a value of 10
		fmt.Println("Executing the if block")
	}

	//Example of if..else if ..else if..else if..else
	counter :=300

	if (counter == 0) {
		fmt.Println("Counter value is 0")
	} else if counter == 1 {
		fmt.Println("Counter value is 1")
	} else if counter == 2 {
		fmt.Println("Counter value is 2")
	} else {
		fmt.Println("Counter value is not 0, 1, or 2 and the value of the counter is : ",counter)
	}

	//Use of logical operator &&
	if(counter<400 && counter>301) {
		fmt.Println("Counter value is within a range")
	}
}