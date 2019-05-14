package main

import "fmt"

type employee_data struct {
	name string
	id string
	office string
}

type employee_interface interface {
	getOffice() string
	getName() string
}

func getEmployeeData(employee employee_interface) {
	employee.getOffice()
	employee.getName()
}

func main() {
	employee:= employee_data {
		name:"Nikhil",
		id:"9012",
		office:"xyz",
	}
	fmt.Printf("%T\n",employee)
	fmt.Printf("%T\n",&employee)

	//getEmployeeData(employee)
	getEmployeeData(&employee)

}

//attach get office to struct - employee data and this is nonpointer receiver
func (e employee_data) getOffice() string {
	fmt.Println("The office is ",e.office)
	return e.office
}

//attach get name to struct - employee data and this is pointer receiver
func (e *employee_data) getName() string {
	fmt.Println("The name is ",e.name)
	return e.name
}
