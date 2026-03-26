package endpoint

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"ToDo/utilities/storage"
)

func DoneHandler(w http.ResponseWriter, r *http.Request) {
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
        http.Error(w, fmt.Sprintf("Invalid ID: %v", err), http.StatusBadRequest)
        return
    }
	
	storageFile, err := storage.StorageFilePath()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting storage path: %v", err), http.StatusInternalServerError)
		return
	}

	readjson, err := os.ReadFile(storageFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(readjson, &tsk)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error unmarshaling file: %v", err), http.StatusInternalServerError)
		return
	}

	for i := range tsk {
    if tsk[i].ID == int(id) {
        tsk[i].Mark = true  
            break  
        }
    }

	updateJson, err := json.MarshalIndent(tsk, "", "  ")
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to marshal: %v", err), http.StatusInternalServerError)
        return
    }

    err = os.WriteFile(storageFile, updateJson, 0644)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error writing file: %v", err), http.StatusInternalServerError)
        return  
    }

	w.WriteHeader(http.StatusOK) 
}