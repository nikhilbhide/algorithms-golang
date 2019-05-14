package HelloWorld

import "fmt"

var message string = "Welcome to the hello world"
var AppName string = "Hello World"

func displayMessage() string {
	var myVar int = 100
	fmt.Println("Hello World")
	return message
	fmt.Println(myVar)
}

func GetAppName() string {
	return AppName
}

func GetDisplayMessage() string {
	return displayMessage()
}

func tryFileLevelScoping() {
	fmt.Println(AppName)
	fmt.Println(message)
	GetAppName()
	displayMessage()
}