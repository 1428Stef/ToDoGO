package storage

import (
    "os"
    "path/filepath"
)

func StorageFilePath() (string, error) {
    root := os.Getenv("TODO_ROOT")
    if root == "" {
        root = "."
    }
    return filepath.Join(root, "storage", "storage.json"), nil
}