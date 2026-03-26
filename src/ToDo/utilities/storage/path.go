package storage

import (
	"path/filepath"
)


func StorageFilePath() (string, error) {
	return filepath.Join("storage", "storage.json"), nil
}

