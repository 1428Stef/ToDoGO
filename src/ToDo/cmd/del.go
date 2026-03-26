package cmd

import (
	"os"
	"encoding/json"
	"ToDo/utilities/storage"
)

func Del(id int) error { 
	
	storageFile, err := storage.StorageFilePath()
	if err != nil {
		return err
	}

	readjson, err := os.ReadFile(storageFile)
	if err != nil {
		return err
	}

	var tsk []Task

	if err := json.Unmarshal(readjson, &tsk); err != nil {
		return err
	}

	found := false

	for i, t := range tsk {
		if t.ID == id {
			tsk = append(tsk[:i], tsk[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return nil
	}

	newjson, err := json.MarshalIndent(tsk, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(storageFile, newjson, 0644)
}
