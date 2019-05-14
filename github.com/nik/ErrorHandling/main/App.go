package main

import (
	"fmt"
	"github.com/nik/ErrorHandling/errorhandling"
)

func main() {
	//invoke the divide with 1 and 0
	res,err:=errorhandling.Divide(1,0)
	//check the error
	if(err==nil) {
		fmt.Println(res)
	} else {
		//defer function
		defer func() {
			fmt.Println("This is error")
		} ()
		//raise the panic
		panic(err)
	}
}