package main

import "fmt"

type person struct {
	first_name string
	last_name  string
}

type sports_person struct {
	person person
	sports string
}

//define a function (behavior for person) with receiver person
func (p person) speak() {
	fmt.Println("I can speak as I am human")
}

//define a function (behavior for person) with receiver person
func (p sports_person) speak() {
	fmt.Println("Either my cricket bat or ball speaks as I am a cricketer")
}

func main() {
	//define a variable for a person
	p1 := person{
		last_name:  "Kohli",
		first_name: "Virat",
	}
	fmt.Println(p1)
	//invoke speak for a person
	p1.speak()

	//define a variable for a sports_player
	sports_player := sports_person{
		person: p1,
		sports: "Cricket",
	}

	//invoke speak for a sports_player
	sports_player.speak()
	//invoke the parent function by explicitly calling the speak on person
	sports_player.person.speak()
}