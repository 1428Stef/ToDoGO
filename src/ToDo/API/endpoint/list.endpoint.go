package endpoint

import (
	"fmt"
	"net/http"
	"os"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	readjson, err := os.ReadFile("../storage/storage.json")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(readjson))
}