package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

// Hello Capital fn name = to be called by an external fn
// Hello returns a greeting for the named person or error.
func Hello(name string) (string, error) {
	//fn returns 2 values: string AND error

	if name == "" {
		return "", errors.New("no name given")
	}
	// Return a greeting that embeds the name in a message.
	// used to be =>  message := fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomFormat(), name)

	//message := fmt.Sprint(randomFormat()) TODO - to see what a failing test looks like, uncomment this

	//nil = no error, indicating a success return
	return message, nil
}

// Hellos returns a map that associates ea of the named ppl with a greeting msg
func Hellos(names []string) (map[string]string, error) {
	//given an array of names
	//return a map where keys are strings and key values are strings

	//declare a map, named messages to assoc. names with messages
	messages := make(map[string]string)

	//for each name of all names, call Hello fn on the name
	for _, name := range names {
		message, err := Hello(name)
		//if Hello fn doesn't return err of nil, return nil and err msg
		if err != nil {
			return nil, err
		}
		//else, set messages[name] = to the msg
		messages[name] = message
	}
	//return the map of messages, with err of nil
	return messages, nil
}

// lowercase function name = to be called within same function file
// randomFormat returns a greeting in a random format
func randomFormat() string {
	//formats = [1, 2, 3]
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v.",
		"Hola, %v.",
	}
	//rand.Intn(num) -- returns a random num between (0,n)
	msgnum := rand.Intn(len(formats))
	log.Println("the line", formats[msgnum])

	return formats[msgnum]
}
