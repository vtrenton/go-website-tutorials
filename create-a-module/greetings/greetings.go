package greetings

import (
	"fmt"
	"errors"
	"math/rand"
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
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages.
// The returned message is selected at random
func randomFormat() string {
	// I know this tutorial wants me to use a slice but i'm gonna make an Array
	// I know the total number of messages and likely this Array wont grow much
	formats := [...]string{
		"Hi, %v. Welcome!",
		"Great to see you %v.",
		"Hello %v!",
	}
	// Return a randomly selected message format
	// by specifying a random index for the array of Formats
	return formats[rand.Intn(len(formats))]
}