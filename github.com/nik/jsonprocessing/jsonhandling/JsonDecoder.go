package jsonhandling

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Decode() {
	var personInstance person
	rdrInstance := strings.NewReader(`{"Firstname":"Mahendrasingh","Lastname":"Dhoni","Age":37}`)
	err := json.NewDecoder(rdrInstance).Decode(&personInstance)
	if (err == nil) {
		fmt.Println(personInstance.Firstname)
		fmt.Println(personInstance.Lastname)
		fmt.Println(personInstance.Age)
	}
}