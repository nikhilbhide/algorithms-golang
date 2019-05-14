package main
import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	standardFields := log.Fields{
		"hostname": "staging-1",
		"appname":  "foo-app",
		"session":  "1ce3f6v",
	}

	log.WithFields(standardFields).WithFields(log.Fields{"string": "foo", "int": 1, "float": 1.1}).Info("My first ssl event from Golang")

}
