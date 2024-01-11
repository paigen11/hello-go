package main

import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	// set properties of predefined Logger
	// including the log entry prefix and flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// slice of names
	names := []string{"Bill", "Charlie", "Denise"}

	// request greeting messages for names
	messages, err := greetings.Hellos(names)
	// if error returned, print it to the console and exit program
	if err != nil {
		log.Fatal(err)
	}

	// if no error returned, print greeting message to console
	fmt.Println(messages)
}
