package services

import (
	"os"
)

func SaveToFile(data []byte, filePath string) error {

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}
