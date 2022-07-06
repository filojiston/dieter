package main

import (
	"dieter/commands/add"
	"dieter/commands/get"
	"flag"
	"fmt"
	"os"
)

func main() {
	addSet := flag.NewFlagSet("add", flag.ExitOnError)
	getSet := flag.NewFlagSet("get", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("you must use add or get command.")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]
	switch command {
	case "add":
		add.Handle(addSet, args)
		return
	case "get":
		get.Handle(getSet, args)
		return
	}
}
