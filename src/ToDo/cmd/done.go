package cmd

import (
	"encoding/json"
	"os"
)

func Done(id int) error {
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
			tsk[i].Mark = true
			
			updateJson, err := json.MarshalIndent(tsk, "", " " )
			if err != nil {
				return err
			}
			err = os.WriteFile("storage/storage.json", updateJson, 0644)
			if err != nil {
				return err
			}
			return nil

		}
	}

	return err
}