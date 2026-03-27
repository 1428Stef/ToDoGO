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
			fmt.Fprintln(os.Stderr, "error: --title is required")
			os.Exit(1)
		}
		
		if err := cmd.Add(*title); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "list":
		if err := cmd.List(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "done":
		doneCmd := flag.NewFlagSet("done",flag.ExitOnError)
		id := doneCmd.Int("id", 0, "")
		doneCmd.Parse(os.Args[2:]) 
		if *id == 0 {
			fmt.Fprintln(os.Stderr, "error: --id is required")
			os.Exit(1)
		}

		if err := cmd.Done(*id); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "del":
		delCmd := flag.NewFlagSet("del", flag.ExitOnError)
		id := delCmd.Int("id", 0, "")
		delCmd.Parse(os.Args[2:])
		if *id == 0 {
			fmt.Fprintln(os.Stderr, "error: --id is required")
			os.Exit(1)
		}

		if err := cmd.Del(*id); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "help":
		cmd.Help()
	case "edit":
		editCmd := flag.NewFlagSet("edit", flag.ExitOnError)
		id := editCmd.Int("id", 0, "")
		title := editCmd.String("title", "", "")
		editCmd.Parse(os.Args[2:])
		if *id == 0 {
			fmt.Fprintln(os.Stderr, "error: --id is required")
			os.Exit(1)
		}
		if *title == "" {
			fmt.Fprintln(os.Stderr, "error: --title is required")
			os.Exit(1)
		}
		if err := cmd.Edit(*id, *title); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %s \n", os.Args[1])
		os.Exit(1)
	}
}
