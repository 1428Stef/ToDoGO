package cmd

import (
	"ToDo/utilities/id"
	"ToDo/utilities/date"
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

	readjson, err := os.ReadFile("storage/storage.json")
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

	if err := os.WriteFile("storage/storage.json", jsonMar, 0644); err != nil {
		return err
	}

	fmt.Printf("New task created: %s\n", titleTask)
	return nil
}