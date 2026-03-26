package endpoint

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func DelHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to process your request: %v", err), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var tsk []Task

	idStr := string(body)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID: %v", err), http.StatusInternalServerError)
	}

	readjson, err := os.ReadFile("../storage/storage.json")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(readjson, &tsk)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error unmarshaling file: %v", err), http.StatusInternalServerError)
		return
	}

	found := false
	for i, t := range tsk {
		if t.ID == int(id) {
			tsk = append(tsk[:i], tsk[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		http.Error(w, fmt.Sprintf("Task not found: %v", err), http.StatusInternalServerError)
		return
	}

	updateJson, err := json.MarshalIndent(tsk, "", "  ")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal: %v", err), http.StatusInternalServerError)
	}

	os.WriteFile("../storage/storage.json", updateJson, 0644)
}