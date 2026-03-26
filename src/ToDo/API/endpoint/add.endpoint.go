package endpoint

import (
	"ToDo/utilities/date"
	"ToDo/utilities/id"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Task struct {
	Title string `json:"title"`
	Mark bool `json:"mark"`
	Date string `json:"date"`
	ID int `json:"id"`
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to process your request: %v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	t := Task {
		Title: string(body),
		Mark: false,
		Date: date.GetTime(),
		ID: id.Generate(),
	}
	var tsk []Task 
	
	readjson, err := os.ReadFile("../storage/storage.json")
	if err == nil && len(readjson) > 0 {
		if err := json.Unmarshal(readjson, &tsk); err != nil {
			http.Error(w, fmt.Sprintf("Error unmarshaling file: %v", err), http.StatusInternalServerError)
			return
		}
	}
	tsk = append(tsk, t)

	jsonMar, err := json.MarshalIndent(tsk, "", "  ")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling file: %v", err), http.StatusInternalServerError)
		return
	}

	if err := os.WriteFile("../storage/storage.json", jsonMar, 0644); err != nil {
		http.Error(w, fmt.Sprintf("Error writing file: %v", err), http.StatusInternalServerError)
		return
	}
}