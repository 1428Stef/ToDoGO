package cmd

import (
	"encoding/json"
	"fmt"
	"ToDo/utilities/storage"
	"os"
)

func List() error {
	storageFile, err := storage.StorageFilePath()
	if err != nil {
		return err
	}

	readjson, err := os.ReadFile(storageFile)
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