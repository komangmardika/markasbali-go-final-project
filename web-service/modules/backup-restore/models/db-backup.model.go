package models

import (
	"gorm.io/gorm"
	"markasbali_go_final_project/web-service/modules/bases/models"
)

type DbBackup struct {
	models.Model
	FileName string `json:"file_name"`
	DbId     uint   `json:"db_id"`
	DbRef    Db     `gorm:"foreignKey:DbId"`
}

func (db *DbBackup) Create(conn *gorm.DB) error {
	return conn.Model(DbBackup{}).
		Create(&db).Error
}

func (db *DbBackup) GetAll(conn *gorm.DB) ([]DbBackup, error) {
	var dbs []DbBackup
	return dbs, conn.Find(&dbs).Error
}

func (db *DbBackup) GetById(conn *gorm.DB) (DbBackup, error) {
	return DbBackup{}, conn.Model(DbBackup{}).
		Where("id = ?", db.ID).
		Where("deleted_at IS NOT NULL").
		Take(&db).Error
}
