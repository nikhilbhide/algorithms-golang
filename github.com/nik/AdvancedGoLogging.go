package main

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

// To use lumberjack with the standard library's log package, just pass it into
// the SetOutput function when your applic ation starts.
func main() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "log_dir/logger.log",
		MaxSize:    1, // megabytes
		//MaxBackups: 100,
		MaxAge:     10,   // days
		Compress:   true, // disabled by default
	})

	for i := 0; i < 10000000; i++ {

		log.Print("This is my logger : ", i)
	}
}
