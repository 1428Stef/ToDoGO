package cmd

import (
	"encoding/json"
	"os"
	"ToDo/utilities/storage"
)

func Edit(id int, newTitle string) error {
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

	for i := range tsk {
		if tsk[i].ID == id {
			tsk[i].Title = newTitle
			
			updateJson, err := json.MarshalIndent(tsk, "", " ")
			if err != nil {
				return err
			}
			err = os.WriteFile(storageFile, updateJson, 0644)
			if err != nil {
				return nil 
			}
			return nil
		}
	}
	
	return err
}