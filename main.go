package main

import (
	"Go_Interpreter/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s. This is the Monkey Programming Language by Thornsten Ball.\n",
		user.Username)
	fmt.Printf("Type in commands to test the interpreter\n")
	repl.Start(os.Stdin, os.Stdout)
}
