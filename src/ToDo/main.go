package main

import (
	"ToDo/cmd"
	"fmt"
	"os"
)

func main() {
	menu := os.Args; if len(menu) < 2 {
		fmt.Println("Please enter 'help' for all commands.")
		return
	}

	if menu[1] == "help" {
		cmd.Help()
		return
	}

	switch menu[1] {
	case "add":
		cmd.Add(menu[2])
	default:
		fmt.Printf("Unknown command: %s \n", menu[1])
		os.Exit(1)
	}

}
