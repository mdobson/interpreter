package main

import (
	"fmt"
	"os"
	"os/user"

	"example.com/interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! REPL v0.0.1 of UnknownScript\n", user.Username)
	fmt.Printf(">>Use exit or Ctrl-C to quit<<\n")
	fmt.Printf(">>Ready to receive commands<<\n")
	repl.Start(os.Stdin, os.Stdout)
}
