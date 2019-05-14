package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	x:=10
	y:=10
	c:= make(chan int)
	for i:=0; i<x ; i++ {
		go func() {
			for j := 0; j < y; j++ {
				c <- j
			}
		}()
	}
	//close(c)

	/*for k:=0 ; k<x*y; k++ {

		println(k, <-c)
	}*/

	/*for value := range c {
		println(value)

	}
*/
//	var ans1, ans2 int

	/*fmt.Print("Name: ")
	nim, err := fmt.Scan(&ans1, &ans2)
	if(err==nil) {
		fmt.Print("entered value is ", ans1,nim)
	} else {
		panic(err)
	}*/

	logfile,err:= os.Create("log.txt")
	log.SetOutput(logfile)
	defer logfile.Close()

	f, err:=os.OpenFile("names1.txt",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	if(err==nil) {
data := [] byte ("Weitisklj")
	f.Write(data)
	}
	f.Close()

	o,error := os.Open("names0as0.txt")
	log.Fatal(error)

	bytes, error := ioutil.ReadAll(o)
	if(error==nil) {
		fmt.Println(string(bytes))

	}
	o.Close()
}