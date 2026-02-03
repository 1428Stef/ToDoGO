package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

func show() error {
	data, err := os.ReadFile("storage/storage.json")
	if err != nil {
		return err
	}

	var tasks string
	err = json.Unmarshal([]byte(data), &tasks)
	if err != nil {
		return err
	}

	fmt.Println(tasks)
	return err
}