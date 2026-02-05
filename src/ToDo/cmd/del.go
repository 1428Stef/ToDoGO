package cmd

import (
	"os"
	"encoding/json"
)

func Del(id int) error { 
	
	readjson, err := os.ReadFile("storage/storage.json")
	if err != nil {
		return err
	}

	var tasks []Task

	if err := json.Unmarshal(readjson, &tasks); err != nil {
		return err
	}

	found := false

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return nil
	}

	newjson, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("storage/storage.json", newjson, 0644)
}
