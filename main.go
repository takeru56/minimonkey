package main

import (
	"fmt"
	"github.com/takeru56/minimonkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is minimonkey language.\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
