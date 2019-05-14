package main

import "fmt"

type human interface {
	work()
}

type base_person struct {
	first_name string
	last_name string
}

func (base_person) work() {
	fmt.Println("I do not work anywhere because I am a person ...")
}

type player struct {
	person base_person
	sports string
}

type employee struct {
	person base_person
	office_name string
	service_type string
}

func (player_instance player) work() {
	fmt.Println("I play the game : ",player_instance.sports)
}

func (employee_instace employee) work() {
	fmt.Printf("I am a service person and works for %s and in the sector %s \n",employee_instace.office_name,employee_instace.service_type)
}

func give_profession(human_instance human) {
	human_instance.work()
}


func main() {
	person_instance_1 := base_person {
		first_name: "Mahendra Singh",
		last_name: "Dhoni",
	}

	person_instance_2 := base_person {
		first_name: "Chinmay",
		last_name: "Jog",
	}

	player_instance:= player{
		person:person_instance_1,
		sports:"Cricket",
	}

	employee_instance:= employee{
		person:person_instance_2,
		office_name:"XYZ",
		service_type:"HighTech",
	}

	give_profession(player_instance)
	give_profession(employee_instance)
	give_profession(person_instance_1)
	give_profession(person_instance_2)
}
