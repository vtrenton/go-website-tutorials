package main

// import my greetings module
// since I have this stored on github - go mod tidy will pull it from there
// else I could use a replacement function:
// replace example.com/greetings => ../greetings
// interesting since my reference is to what is on github - i need to push often.
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
