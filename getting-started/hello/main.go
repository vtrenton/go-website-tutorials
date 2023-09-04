package main

// fmt is part of the std library - go builtins: https://pkg.go.dev/std
// i can import a package to do some of the work for me
// in this case "rsc.io/quote" to say a statement
// since rsc.io/quote is not from the std library - we will need to pull it with go mod tidy
import ( 
	"fmt"
	"rsc.io/quote"
)

func main() {
	// instead of writing my own Hello World program
	fmt.Println(quote.Go())
}
