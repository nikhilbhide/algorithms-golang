package main

import "fmt"

func main() {

	type person struct {
		first_name string
		last_name string
	}
	p1:= person {
		first_name:"Janes",
		last_name:"Messi",
	}

	type player struct {
		person person
		game  string
	}

	//embedding one struct into another
	sports_player:= player {
		person: person {
			first_name:"Lionel",
			last_name: "Messi",
		},
		game: "football",
	}

	sports_player1:= player {
		person: p1,
		game: "football",
	}

	//declaration of variables having same data type (i.e string)
	//variables are declared on one line and separated by a column
	type government_employee struct {
		person person
		government_office, id string
	}

	//anoonymous struct
	chinese_person:= struct {
		first_name string
		last_name string
	} {
		first_name:"Chang Po Cho",
		last_name:"Po Po",
	}

	fmt.Println(chinese_person)

	fmt.Println(sports_player)
	fmt.Println(sports_player1)


}
