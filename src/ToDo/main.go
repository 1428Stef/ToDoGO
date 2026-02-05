package main

import (
	"ToDo/cmd"
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please enter 'help' for all commands.")
		return
	}

	switch os.Args[1] {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		title := addCmd.String("title", "", "")
		addCmd.Parse(os.Args[2:])
		if *title == "" {
			os.Exit(1)
		}
		
		cmd.Add(*title)
	case "list":
		cmd.List()
	case "done":
		doneCmd := flag.NewFlagSet("done",flag.ExitOnError)
		id := doneCmd.Int("id", 0, "")
		doneCmd.Parse(os.Args[2:]) 
		if *id == 0 {
			os.Exit(1)
		}

		cmd.Done(*id)
	case "del":
		delCmd := flag.NewFlagSet("del", flag.ExitOnError)
		id := delCmd.Int("id", 0, "")
		delCmd.Parse(os.Args[2:])
		if *id == 0 {
			os.Exit(1)
		}

		cmd.Del(*id)
	case "help":
		cmd.Help()
	default:
		fmt.Printf("Unknown command: %s \n", os.Args[1])
		os.Exit(1)
	}
}
