package main

import "fmt"

//func (receiver) function_name (input parameters) (return parameters)
func main() {
	//invoke a simple function
	hello_world()

	x := 10
	y := 20
	sum_result := sum(x, y) //x and y are known as arguments
	fmt.Println("Sum of two numbers x and y : ", sum_result)

	//invoke a function that returns two integers
	first, second := sort_numbers(10, 90)
	fmt.Println(first, second)

	//invoke a function that takes variadic parameter
	//slice_coll:=[]int {1,304,120,20}
	calculate_grp_sum(1, 304, 120, 20)

	fmt.Println("Now invoking the variadic function with nothing - nil")
	//invoke a function that takes variadic parameter
	calculate_grp_sum()

	fmt.Println("Now invoking the variadic function with actual slice variable")
	//invoke a function that takes slice variable
	var slice_x [] int
	fmt.Println(len(slice_x))
	fmt.Println(cap(slice_x))
	slice_x = append(slice_x,10,30,50,80,90)
	calculate_grp_sum(slice_x...)

	//invoking the function having regular input parameter (string in this case) and variadic parameter as well
	calculate_grp_sum_with_behavior("sum",3,40,506)
	//invoking the function having regular input parameter (string in this case) and no variadic parameter as well
	calculate_grp_sum_with_behavior("sum")

	//example of anonymous function
	func (x int) {
		fmt.Println("The value of x is : ",x)
	} (100)


	//example of anonymous function that returns the value
	value := func (x int) (int) {
		fmt.Println("The value of x is : ",x)
		return -1
	} (100)
	fmt.Println("The value is :",value)

	//store function into the function type variable
	func_variable := func () {
		fmt.Println("This function was executed")
	}
	fmt.Printf("%T\n",func_variable)

	//invoking the anonymous function which is stored in the function type variable
	func_variable()

	//store function into the function type variable
	complex_func_variable := func (x int, value string) (string) {
		fmt.Println("The value of x is : ",x)
		fmt.Println("The value of value is : ",value)

		return "Executed!!"
	}

	output:=complex_func_variable(80,"Human")
	fmt.Println("Output is :",output)

	//invoking the anonymous function which is stored in the function type variable
	func_variable()

	defer_demo()
}

//define a simple function
func hello_world() {
	fmt.Println("Welcome to the GoLang world")
}

//define a function sum with two input parameters x and y and return sum (x+y)
func sum(x int, y int) (int) {
	sum := x + y
	return sum
}

//define a function that takes two parameters and returns two values
func sort_numbers(x int, y int) (int, int) {
	var first_number int
	var second_number int

	//compare numbers
	if (x > y) {
		first_number = x
		second_number = y
	} else {
		first_number = y
		second_number = x
	}

	//return two integers
	return first_number, second_number
}

//function that takes variadic parameter
func calculate_grp_sum(numbers ...int) {
	fmt.Print("Displaying the details of the slice\n")
	fmt.Printf("%T\n", numbers)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	var sum int
	//calculate sum by iterating over a slice
	for _, value := range numbers {
		fmt.Println(value)
		sum += value
	}
	fmt.Println("sum of numbers in a slice is : ", sum)
}

//function that takes variadic parameter and other regular parameters
func calculate_grp_sum_with_behavior(behavior string,numbers ...int) {
	fmt.Print("Displaying the details of the slice\n")
	fmt.Printf("%T\n", numbers)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	var sum int
	//calculate sum by iterating over a slice
	for _, value := range numbers {
		fmt.Println(value)
		sum += value
	}
	fmt.Println("behavior of this function is : ", behavior)
	fmt.Println("sum of numbers in a slice is : ", sum)
}

func defer_demo () {
	defer defer_func()
	fmt.Println("I got executed after defer_func")
	var sum int
	for i:=0;i<100; i++ {
		sum+= i
	}
	fmt.Println("The value of sum is :",sum)
}

func defer_func() {
	fmt.Println("I got executed before defer_demo")
}