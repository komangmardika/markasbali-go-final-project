package services

import (
	"os"
	"path/filepath"
)

func RemoveFilesInTmpFolder() error {
	// Folder path
	folderPath := os.Getenv("TMP_FOLDER_PATH")

	// List all files in the folder
	files, err := filepath.Glob(filepath.Join(folderPath, "*"))
	if err != nil {
		return err
	}

	// Delete each file
	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			return err
		}
	}

	return nil
}
