package main

import "fmt"

func main() {
	channel := make(<-chan  int,1)
//	channel<-2
	//fmt.Println(<-channel)
	fmt.Printf("%T\n",channel)
}
