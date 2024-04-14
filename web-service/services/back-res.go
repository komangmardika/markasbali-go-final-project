package services

import (
	"markasbali_go_final_project/web-service/configs"
	"markasbali_go_final_project/web-service/models"
	"os"
)

func GetLatestBackedUpDatabaseList() ([]models.DataListDto, error) {
	db := models.Db{}
	dbs, err := db.GetAllWithLatestBackup(configs.Mysql.DB)

	if err != nil {
		return []models.DataListDto{}, err
	}

	var head []models.DataListDto

	for _, value := range dbs {

		details := models.LatestBackup{
			ID:        value.DbBackup[0].ID,
			FileName:  value.DbBackup[0].FileName,
			Timestamp: value.DbBackup[0].UpdatedAt,
		}

		head = append(head, models.DataListDto{
			DatabaseName: value.DatabaseName,
			LatestBackup: details,
		})

	}

	return head, nil
}

func GetOneDatabaseWithHistory(dbName string) (models.DataSingleDto, error) {
	db := models.Db{
		DatabaseName: dbName,
	}
	err := db.GetOneDatabaseWithAllBackup(configs.Mysql.DB)
	if err != nil {
		return models.DataSingleDto{}, err
	}
	var details []models.LatestBackup
	for _, value := range db.DbBackup {
		details = append(details, models.LatestBackup{
			ID:        value.ID,
			FileName:  value.FileName,
			Timestamp: value.UpdatedAt,
		})
	}

	head := models.DataSingleDto{
		DatabaseName: db.DatabaseName,
		Histories:    details,
	}
	return head, err
}

func GetDownloadLatestBackedUpByDatabase(id uint) ([]byte, error) {

	fileContent, err := os.ReadFile("example.txt")
	if err != nil {
		return []byte{}, err
	}

	return fileContent, nil
}
