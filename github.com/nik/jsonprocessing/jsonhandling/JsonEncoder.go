package jsonhandling

import (
	"encoding/json"
	"os"
)



func Encode() {
	personInstance:= person {Firstname:"Virat",Lastname:"Kohli",Age:31}
	json.NewEncoder(os.Stdout).Encode(personInstance)
}


