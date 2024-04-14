package services

import (
	"fmt"
	"os"
)

func CheckOrCreateFolder(dbName string) (string, error) {
	folderPath := os.Getenv("STORAGE_FOLDER_PATH") + dbName
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// Folder doesn't exist, create it
		err := os.MkdirAll(folderPath, 0755)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return "", err
		}
		return folderPath, nil
	} else if err != nil {
		fmt.Println("Error checking folder existence:", err)
		return "", err
	}

	return folderPath, nil
}
