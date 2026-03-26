package endpoint

import (
	"fmt"
	"net/http"
	"os"
	"ToDo/utilities/storage"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(readjson))
}