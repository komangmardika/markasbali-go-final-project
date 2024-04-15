package services_test

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"markasbali_go_final_project/cli-service/models"
	"markasbali_go_final_project/cli-service/services"
	"os"
	"testing"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("cannot find .env file")
	}
}

func TestRequestLatestBackupInfo(t *testing.T) {
	resp, err := services.RequestRestoreToServer("restore", "all-db-latest-history", "pt_abc")
	assert.Nil(t, err)
	var detail []models.DataListDto
	err = json.Unmarshal([]byte(resp), &detail)
	fmt.Println(resp)
	fmt.Println(detail)
}

func TestDownloadFile(t *testing.T) {
	Init()
	resp, err := services.RequestFileToServer(19)

	assert.Nil(t, err)
	err = services.SaveToFile(resp, "../"+os.Getenv("TMP_FOLDER_PATH")+"foo.zip")
}

func TestUnzipFile(t *testing.T) {
	Init()
	f := models.MySqlConnWithBackup{
		SqlFileName: "mysql-2024-04-13-01-11-21-pt_xyz-3b46e3e8-47d2-493b-9764-26864096f6f0.sql",
		FileName:    "mysql-2024-04-13-01-11-21-pt_xyz-3b46e3e8-47d2-493b-9764-26864096f6f0.zip",
		TmpFolder:   "../tmp/",
	}

	_, err := services.UnzipFile(f)

	assert.Nil(t, err)
}
