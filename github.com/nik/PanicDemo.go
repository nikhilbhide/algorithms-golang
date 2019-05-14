package main

import (
	"fmt"
	"os"
)

func main() {
	//inherent error handling in golang
	//this includes the population of err variable of type error and then checking whether err not nil

	file_handler, _:= os.Open("employee.json")
	defer func(file_handler *os.File) {
		//close any open resources
		if(file_handler!=nil) {
			file_handler.Close()
		}

		//check whether recover is not nil
		if r:=recover(); r!=nil {
			//handle the panicking sequence
			fmt.Println("The value returned from the panic sequence : ",r)
			fmt.Printf("%T\n",r)
		}
	} (file_handler)

	//invoke display function
	display(0)
}

//function that displays counter value until value is greater than or equal to 4
//it raises panic in case value of counter is equal to 4
func display(counter int) {
	defer func (counter int) {
		fmt.Println("Executing deferred function in display", counter )
	} (counter)

	if(counter>=4) {
		panic(fmt.Sprintf("%v",counter))
	}
	display(counter+1) //display(5) will never be called
}
