package services_test

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"markasbali_go_final_project/web-service/configs"
	"markasbali_go_final_project/web-service/modules/backup-restore/models"
	"testing"
)

func Init() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		fmt.Println("env not found, using system env")
	}
	configs.OpenDB(false)
}

func TestGetLatestBackedUpDatabaseList(t *testing.T) {
	Init()
	db := models.Db{}
	dbs, err := db.GetAllWithLatestBackup(configs.Mysql.DB)

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
	jsonString, err := json.Marshal(head)
	fmt.Println(string(jsonString))
	assert.Nil(t, err)
}

func TestGetOneDatabaseWithAllBackup(t *testing.T) {
	Init()
	db := models.Db{
		DatabaseName: "pt_abc",
	}
	err := db.GetOneDatabaseWithAllBackup(configs.Mysql.DB)
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
	jsonString, err := json.Marshal(head)
	fmt.Println(string(jsonString))
	assert.Nil(t, err)
}
