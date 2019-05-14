package main

import "fmt"

func main() {
	//for loop structure ://init     //condition //post statement

	//Iterate over 100 times
	for counter := 0; counter < 100; counter = counter + 2 {
		fmt.Println("This is for loop iteration no : ", counter)
	}

	//Iterate over 50 times by ensuring that you increment counter by 2
	for counter := 0; counter < 100; counter = counter + 2 {
		fmt.Println("This is for loop iteration no : ", counter)
	}

	//Example of nested looping
	for outerCounter := 0; outerCounter < 100; outerCounter++ {
		for innerCounter := 0; innerCounter < 100; innerCounter++ {
			fmt.Println("The values of outer counter and inner counter are : ", outerCounter, innerCounter)
		}
	}

	//This is the example of using var variable in the for loop
	var counter = 1
	for counter < 10 {
		fmt.Println("The value of counter is : ", counter)
		counter = counter * 2
	}
	//The var - counter variable is accessible even after the looping is over
	fmt.Println("The value of counter after all iterations are over is : ", counter)

	//Use of break keyword in the loop. Stop after counter value becomes 52
	for counter := 0; counter < 100; counter = counter + 2 {
		fmt.Println("This is for loop iteration no : ", counter)
		if counter > 50 {
			fmt.Println("The break condition has been satisfied so breaking out")
			break
		}
	}

	//Use of continue keyword in the loop. Continue in case of the odd number
	for counter := 1; counter <= 100; counter++ {
		if counter%2 != 0 {
			//Odd number case
			continue
		}
		//Its even number so display the value of the number
		fmt.Println("This is the even number and value is : ", counter)
	}

	/*//Tweak for statement to make use of for as while
	for ; 1==1; {  //equivalent to while (true) { }
		fmt.Println("Iterative forever")
	}

	//Tweak for statement to make use of for as while
	for 1==1 {  //equivalent to while (true) { }
		fmt.Println("Iterative forever")
	}
	*/
}
