package main

import "fmt"

//Declare a constant iota with 3 values 0,1,& 2 and associated variables are a,b,& c
const (
	a = iota
	b  = iota
	c = iota
)

//Declare a constant iota with 3 values 0,1,& 2 and associated variables are a,b,& c
const (
	d = iota * 10
	e  = iota * 10
	f = iota * 10
)

//Declare a custom data type of base type int
type Day int

func main() {
	//Display the values of iota constants a, b, & c
	fmt.Println("Values of iota constatns a, b, c are : ", a, b, c)
	//Display the data type of iota constatnt a
	fmt.Printf("Data type of iota constant a is : %T\n", a)

	//Display the values of iota constants d, e, & f
	fmt.Println("Values of iota constatns d, e, f are : ", d, e, f)
	//Display the data type of iota constatnt e
	fmt.Printf("Data type of iota constant a is : %T\n", e)


	//Delcare constant iota for Day
	const (
		Sunday Day = iota + 1
		Monday Day = iota + 1
		Tuesday Day = iota + 1
		Wednesday Day = iota + 1
		Thursday Day = iota + 1
		Friday Day = iota + 1
 		Saturday Day = iota + 1
	)

	//Display the output of a query related to iota Day
	fmt.Println("Whether Sunday is before Monday? : ", Sunday<Monday)
	fmt.Println("The value of Monday is : ",Monday)
	fmt.Println("The value of Saturday is : ",Saturday)

	//Delcare constant iota for auto incremental feature of golang
	const (
		h  = iota //h is 0
		i = -200 //Unused iota. This implies that value 1 is never used
		j  = iota //value over here is 2
	)

	//Display the values of h, i, & j
	fmt.Println("Values of iota constatns h, i, j are : ", h, i, j)
}