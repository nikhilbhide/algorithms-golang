package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

//function that divides first number by second number
func divide(num_1 float32, num_2 float32) (float32, error) {
	if num_2 == 0 {
		return 0, errors.New("Not possible to divide by zero")
	}

	return num_1 / num_2, nil
}

func main() {

	var num_1 float32 = 10
	var num_2 float32 = 0

	//invoke divide function
	ret, err :=  divide(num_1,num_2)

	//handle error
	if err != nil{
		log.Print(err)
	}

	fmt.Println(ret)

	//direct log files to log file
	logger, err := os.OpenFile("sample_log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer logger.Close()
	log.Print("The value of divide operation is : ",ret)
	log.Print("The error of divide operation is : ",err)

	log.SetOutput(logger)

	defer func () {
		fmt.Println("This is to demonstrate log.fatal and log.panic")
	} ()

	log.Panicln(err)


	//fmt.Println("is",ret)
}