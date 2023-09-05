package greetings

import (
	"fmt"
	"errors"
)

// Hello returns a greeting for the named person
// note that the func name starts with a capitol letter
// this allows it to be accessed outside of the current package
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == ""{
		return "", errors.New("Empty name")
	}

	// If a name is given
	// return a greeting that embeds the name in a greeting message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
