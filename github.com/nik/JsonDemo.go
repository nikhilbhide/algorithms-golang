package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//define struct team
type team struct {
	Name, Country, Sports string
	Players [] string
}

func marshal() {
	//instantiate team instance
	team_india := &team{
		"India_WC_2019",
		"India",
		"Cricket",
		[] string{"Virat", "Dhoni", "Dhawan", "Pant"},
	}

	team_aus := &team{
		"Aus_WC_2019",
		"Austrailia",
		"Cricket",
		[] string{"Paine", "Warner", "Smith"},
	}

	teams:= [] *team{team_india,team_aus}

	//display slice of two records
	fmt.Println(teams)

	//marshal team instance to json data
	json_data, error := json.Marshal(teams)
	if (error == nil) {
		//check for error and then display the json data
		fmt.Println(string(json_data))
	}
}

func unmarshal() {
	//declare the map of string and empty value
	var result_data map[string]interface{}

	// Open jsonFile employees.json (it should be in your classpath)
	jsonFile, err := os.Open("employees.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(jsonFile)
	}

	//use defer to close the file anyhow
	defer jsonFile.Close()
	//read the json file into the byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Printf("%T\n",byteValue)
	//unmarshal the bytes and store the result into struct format in result_data
	json.Unmarshal([]byte(byteValue), &result_data)
	fmt.Println(result_data["users"])
}

func main() {
	marshal()
	unmarshal()
}
