package models

import (
	"gorm.io/gorm"
)

type DbBackup struct {
	Model
	FileName string `json:"file_name" valid:"required"`
	DbId     uint   `json:"-" gorm:"foreignKey:DbId" valid:"required"`
}

func (db *DbBackup) Create(conn *gorm.DB) error {
	return conn.Model(DbBackup{}).
		Create(&db).Error
}

func (db *DbBackup) GetById(conn *gorm.DB) (DbBackup, error) {
	return DbBackup{}, conn.Model(DbBackup{}).
		Where("id = ?", db.ID).
		Where("deleted_at IS NOT NULL").
		Take(&db).Error
}
