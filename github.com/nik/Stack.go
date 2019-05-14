package main

import (
	"errors"
	"fmt"
)

var curr_pointer int = -1
var stack []int

func main() {
	stack = make([]int, 10, 20)
	push(1)
	push(2)
	push(3)
	push(4)
	push(5)
	push(6)
	push(7)

	pop()
	pop()
	pop()
	pop()
	pop()
	pop()
	pop()
	element, error := pop()
	fmt.Println("Element : ",element)
	fmt.Println("Error : ", error)

	fmt.Println(stack)
}

func push(element_to_add int) (bool,error) {
	curr_pointer = curr_pointer + 1
	stack[curr_pointer] = element_to_add
	return true,nil
}

func pop() (int,error) {
	fmt.Println("curr_pointer : ",curr_pointer)
	if curr_pointer < 0 {
		fmt.Println("Stack is empty...")
		return -999, errors.New("Stack is empty")
	} else {
		element_to_return := stack[curr_pointer]
		stack[curr_pointer] = 0
		curr_pointer = curr_pointer - 1
		return element_to_return, nil
	}
}
