package utils

import (
	"final-project/kelas-beta-golang/config"
	"final-project/kelas-beta-golang/model"
	"time"
)

func GetList() ([]model.AutoRes, error) {
	var autores model.AutoRes
	return autores.GetAll(config.Mysql.DB)
}

func GetAll() ([]model.AutoRes, error) {
	autores := model.AutoRes{}
	return autores.GetAll(config.Mysql.DB)
}

func GetDistinct() ([]string, error) {
	autores := model.AutoRes{}
	return autores.GetDistinct(config.Mysql.DB)
}

func GetLatestByDBName(nama_db string) (model.AutoRes, error) {
	autores := model.AutoRes{
		Nama_Database: nama_db,
	}
	return autores.GetLatestByDBName(config.Mysql.DB)
}

func InsertData(data model.AutoRes) (*model.AutoRes, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	car, err := data.Create(config.Mysql.DB)
	return car, err
}

