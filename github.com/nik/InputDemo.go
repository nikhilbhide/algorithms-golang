package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inputSingleNumber() {
	//define a variable of type int
	var number int

	fmt.Printf("Enter a number : ")
	//accept the data from the console
	entries, err := fmt.Scanf("%d", &number) //it has to be an integer
	if err==nil {
		fmt.Println("You have entered : ", number)
		fmt.Println("Total number of entries :", entries)
	} else {
		fmt.Println(err)
		panic(err)
	}
}

func inputArrayNumbers() {
	fmt.Printf("Enter a list of numbers : ")
	//use scanner to read the data from console
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	line := scanner.Text()

	//as we have received a text, convert text into array of strings separated by a comma
	splitted_strings := strings.Split(line," ")
	fmt.Println(splitted_strings)

	//declare a slice int variable to hold the integers
	nums := make([] int, len(splitted_strings))
	//try to convert the string to integers one by one.. if there is no parsing issue then store the converted int
	for index, value:=range splitted_strings {
		int_value,err := strconv.Atoi(value)
		if err==nil {
			nums[index] = int_value
		} else {
			fmt.Println("Error occured for value : ",value)
			fmt.Println("Error is : ",err)
		}
	}
	//display the slice of the int
	fmt.Println(nums)
}

func main() {
	inputArrayNumbers()
}
