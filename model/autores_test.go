package model_test

import (
	"encoding/json"
	"final-project/kelas-beta-golang/config"
	"final-project/kelas-beta-golang/model"
	"fmt"

	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init()  {
	err := godotenv.Load("../.env")
	if err != nil{
		fmt.Println("env not found, using global .env")
	}
	config.OpenDB()
}

func TestCreate(t *testing.T)  {
	Init()

	// autoresData := []model.AutoRes{
	// 	{
	// 		Nama_Database: "Inventaris",
	// 		Nama_File_Backup: "mysql-2024-04-05-11-58-00-Inventaris-e7e999ca-9d08-40ca-9763-1427db049d2e.sql.zip",
	// 	},
	// 	{
	// 		Nama_Database: "Keuangan",
	// 		Nama_File_Backup: "mysql-2024-04-05-11-21-00-Keuangan-6b8067ef-ec81-4c04-900d-e959fa6c33ae.sql.zip",
	// 	},
	// 	{
	// 		Nama_Database: "PemesananMakanan",
	// 		Nama_File_Backup: "mysql-2024-04-05-14-15-00-PemesananMakanan-07c70105-3397-40ba-96b7-ba02b4958e0f.sql.zip",
	// 	},
	// 	{
	// 		Nama_Database: "Pendidikan",
	// 		Nama_File_Backup: "mysql-2024-04-05-12-18-48-Pendidikan-49e597-30c1-4e52-92a1-cf8e918499be.sql.zip",
	// 	},
	// 	{
	// 		Nama_Database: "TokoBuku",
	// 		Nama_File_Backup: "mysql-2024-04-05-13-10-TokoBuku-fd37150b-5f94-477a-bb2b-dba45ad66cc3.sql.zip",
	// 	},
	// }

	// autoresData := []model.AutoRes{
	// 	{
	// 		Nama_Database: "Inventaris",
	// 		Nama_File_Backup: "mysql-2024-04-06-11-58-00-Inventaris-7ef2777b-67d3-4019-b6d8-605109a8ca14.sql",
	// 	},
	// 	{
	// 		Nama_Database: "Inventaris",
	// 		Nama_File_Backup: "mysql-2024-04-07-12-58-00-Inventaris-ebf75de9-3438-4eef-a743-eec57637e990.sql",
	// 	},
	// }

	autoresData := []model.AutoRes{
		{
			Nama_Database: "TokoBuku",
			Nama_File_Backup: "mysql-2024-04-05-14-30-TokoBuku-beae61de-dc98-4c67-8ac7-d26e2235caa5.sql",
		},
		{
			Nama_Database: "TokoBuku",
			Nama_File_Backup: "mysql-2024-04-05-14-45-TokoBuku-5d3f8989-c6ed-460f-85f9-4a2eb84ae694.sql",
		},
	}

	// Memasukkan data ke dalam database
    for _, data := range autoresData {
        _, err:= data.Create(config.Mysql.DB)
		
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println("Data berhasil dimasukkan:", data.Nama_Database)

			
		assert.Nil(t, err)
    }
}

func TestGetAll(t *testing.T)  {
	Init()

	autoResData := model.AutoRes{}

	res, err := autoResData.GetAll(config.Mysql.DB)

	assert.Nil(t, err)

	resJson, err := json.Marshal(res)

	if(err != nil){
		log.Fatal(err)
	}

	resJsonString := string(resJson)

	fmt.Println(resJsonString)
}

func TestGetDistinct(t *testing.T)  {
	Init()

	autoResData := model.AutoRes{}

	res, err := autoResData.GetDistinct(config.Mysql.DB)

	assert.Nil(t, err)

	resJson, err := json.Marshal(res)

	if(err != nil){
		log.Fatal(err)
	}

	resJsonString := string(resJson)

	fmt.Println(resJsonString)
}

func TestGetByDBName(t *testing.T)  {
	Init()

	autoResData := model.AutoRes{
		Nama_Database: "Inventaris",
	}

	res, err := autoResData.GetLatestByDBName(config.Mysql.DB)

	assert.Nil(t, err)

	resJson, err := json.Marshal(res)

	if(err != nil){
		log.Fatal(err)
	}

	resJsonString := string(resJson)

	fmt.Println(resJsonString)
}