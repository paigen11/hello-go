package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// create a message using random format
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// hellos returns map that associates each of the named people with greeting message
func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messages
	messages := make(map[string]string)
	// loop through received slice of names, calling Hello function to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in map, associate retrieved message with name
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting message. The returned message is selected at random
func randomFormat() string {
	// a slice of message formats
	formats := []string{
		"Hi %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// return randomly selected message format by specifying random index for slice of formats
	return formats[rand.Intn(len(formats))]
}
