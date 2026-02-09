package cmd

import (
	"encoding/json"
	"os"
)

func Edit(id int, newTitle string) error {
	readjson, err := os.ReadFile("storage/storage.json")
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
			err = os.WriteFile("storage/storage.json", updateJson, 0644)
			if err != nil {
				return nil 
			}
			return nil
		}
	}
	
	return err
}