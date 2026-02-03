package cmd

import (
	"ToDo/id"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Title string
	Mark  bool
	Date  string
	ID int 
}

func Add(titleTask string) error{
	nowTime := time.Now()
	formattedTime := nowTime.Format("2006-01-02 15:04:05")

	t := Task{
		Title: titleTask,
		Mark:  false,
		Date: formattedTime, 
		ID: id.Generate(), 
	}
	var tasks []Task

	data, err := os.ReadFile("storage/storage.json")
	if err == nil && len(data) > 0 {
		if err := json.Unmarshal(data, &tasks); err != nil {
			return err
		}
	}

	tasks = append(tasks, t)

	jsonMar, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
    	return err
	}

	if err := os.WriteFile("storage/storage.json", jsonMar, 0644); err != nil {
    	return err
	}

	fmt.Printf("New task created: %s\n", titleTask)
	return nil
}