package main

import "fmt"

func main() {
	person_map := map[string]int {
		"Nikhil":30,
		"Ihsa":33,
	}

	fmt.Println(person_map["Nikhil"])

	//v,ok:= person_map["Nikhil"]

	//comma-ok idiom
	if v,ok:= person_map["Nikhil"];ok {
		fmt.Println("This is the value",v)
	}

	//adding elements to the map
	person_map["anirudha"]=34

	fmt.Println(person_map)

	//iterate over a loop
	for k,v := range person_map {
		fmt.Println("Key & Value",k,v)
	}

	//delete entries from a map
	delete(person_map,"Nikhil")
	fmt.Println(person_map)

	if v,ok:= person_map["Ihsa"];ok {
		fmt.Println("value exists",v)
		delete(person_map,"Isha")
	}

	person_map["ankit"] = 20
	fmt.Println(person_map)
}