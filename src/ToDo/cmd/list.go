package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

func List() error {
	readjson, err := os.ReadFile("storage/storage.json")
	if err != nil {
		return err
	}
	
	var tsk []Task

	err = json.Unmarshal(readjson, &tsk)
	if err != nil {
		return err
	} 
	
	for _, t := range tsk {
		var mark string
		switch t.Mark {
		case true:
			mark = "✅"
		case false:
			mark = "❌"
		}

		fmt.Printf("Title: %s | date: %s | condition: %s | id: %d \n", t.Title, t.Date, mark, t.ID)
	}

	return err
}