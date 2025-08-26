package main

import (
	"fmt"
	"monkey/repl"
	"os"
)

func main() {
	fmt.Print("monkey REPL\n")
	fmt.Print("Enter a command to get started.\n")

	repl.Start(os.Stdin, os.Stdout)
}
