package main

import (
	"ToDo/cmd"
	"fmt"
	"os"
	"flag"
)

func main() {
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		title := addCmd.String("title", "", "")
		if *title == "" {
			os.Exit(1)
		}


	if len(os.Args) < 2 {
		fmt.Println("Please enter 'help' for all commands.")
		return
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		cmd.Add(*title)
	case "list":
		cmd.List()
	case "help":
		cmd.Help()
	default:
		fmt.Printf("Unknown command: %s \n", os.Args[1])
		os.Exit(1)
	}
}
