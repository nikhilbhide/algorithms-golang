package main

import (
	"fmt"
	"github.com/nik/ErrorHandling/errorhandling"
)

func main() {
	res,err:=errorhandling.Divide(1,0)
	if(err==nil) {
		fmt.Println(res)
	} else {
		defer func() {
			fmt.Println("This is error")
		} ()
		panic(err)
	}
}
