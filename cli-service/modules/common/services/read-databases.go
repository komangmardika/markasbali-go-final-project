package services

import (
	"encoding/json"
	"markasbali_go_final_project/cli-service/modules/common/models"
	"os"
)

func ReadDatabasesJson() ([]models.MySqlConn, error) {

	var configs []models.MySqlConn
	file, err := os.Open("databases.json")

	if err != nil {
		return configs, err
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&configs); err != nil {
		return configs, err
	}

	err = file.Close()

	return configs, nil

}
