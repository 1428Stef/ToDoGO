package cmd

import (
	"ToDo/utilities/id"
	"ToDo/utilities/date"
	"ToDo/utilities/storage"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Title string `json:"title"`
	Mark  bool	`json:"mark"`
	Date  string `json:"date"`
	ID int `json:"id"`
}

func Add(titleTask string) error{
	t := Task{
		Title: titleTask,
		Mark:  false,
		Date: date.GetTime(), 
		ID: id.Generate(), 
	}
	var tsk []Task
	storageFile, err := storage.StorageFilePath()
	if err != nil {
		return err
	}

	readjson, err := os.ReadFile(storageFile)
	if err == nil && len(readjson) > 0 {
		if err := json.Unmarshal(readjson, &tsk); err != nil {
			return err
		}
	}

	tsk= append(tsk, t)

	jsonMar, err := json.MarshalIndent(tsk, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(storageFile, jsonMar, 0644); err != nil {
		return err
	}

	fmt.Printf("New task created: %s\n", titleTask)
	return nil
}