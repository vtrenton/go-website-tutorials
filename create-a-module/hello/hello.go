package main

// import my greetings module
import (
	"fmt"
	"log"
	"github.com/vtrenton/go-website-tutorials/create-a-module/greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it
	message, err := greetings.Hello("")
	// if an error was returned print it to the console and exit
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
