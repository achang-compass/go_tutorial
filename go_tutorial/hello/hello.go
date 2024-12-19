// declare a main package
// a package is way to group functions
// it is made up of all the files in the same directory
package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("log greetings: ") //sets the log entry prefix
	log.SetFlags(0)                  //this disables printing time, source file, and line number

	names := []string{
		"Alice",
		"Vivek",
		"Bobby",
	}

	message, err := greetings.Hellos(names)
	if err != nil {
		//if error returned, prints to console and exit program
		log.Fatal(err)
	}

	//else, prints the message to console
	fmt.Println(message)
}
