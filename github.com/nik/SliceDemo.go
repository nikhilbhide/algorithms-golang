package main

import "fmt"
import "reflect"

func main() {
	//Initiate a slice x of integers
	var array_x [] int
	fmt.Println(array_x)
	fmt.Println(len(array_x))
	fmt.Println(reflect.TypeOf(array_x))
	fmt.Println(append(array_x,1,2,3))
	array_x = append(array_x,1,2,3)
	fmt.Println(array_x)

	x:= []int{24,25,26,27}
	fmt.Println(x)

	//range clause for iteration
	fmt.Println(len(x))
	fmt.Println(x[3])

	for i,v:= range x {
		fmt.Println("index & value",i,v)
	}
	slice:= x[1:3]
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(x[1:3])

	y:= append(array_x,x...)
	fmt.Println(y)
	y = append(y[0:3],y[4:]...)
	fmt.Println(y)

	//create a slice using make - a built in function
	//the first parameter is the type of slice
	//second parameter is length
	//third parameter is the size of underlying array. This much of space is allocated
	//you cna add elements starting from index 0 to index = length-1 (in following case its 9)
	//*** you can not just do this slice_x[99]=3000 in following example
	//in order to do so you need to use append

	slice_x:=make([]int,10,100)
	fmt.Println(len(slice_x))
	fmt.Println(cap(slice_x))
	slice_x[9] =1000

	//if you do not pass second integer argument to make function then capcity = length
	slice_y:=make([]int,10)
	fmt.Println(len(slice_y))
	fmt.Println(cap(slice_y))

	//slice provides dynamic allocation so size is not fixed
	//if length exceeds the capcity then the capacity is increased byh twice
	//folowing append operation will double the size of slice_y
	//internally underlying array is discarded and new array is created at run time
	slice_y = append(slice_y,10)
	fmt.Println(len(slice_y))
	fmt.Println(cap(slice_y))

	lm:= []string{"Lionel","Messi","Football","Argentina","36"}
	vk:= []string{"Virat","Kohli","Cricket","India","34"}

	global_players:= [][]string {lm,vk}
	fmt.Println(global_players)
}
