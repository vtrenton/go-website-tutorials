package main

// import my greetings module
// since I have this stored on github - go mod tidy will pull it from there
// else I could use a replacement function:
// replace example.com/greetings => ../greetings
import (
	"fmt"
	"github.com/vtrenton/go-website-tutorials/create-a-module/greetings"
)

func main() {
	// Get a greeting message and print it
	message := greetings.Hello("Trent")
	fmt.Println(message)
}
