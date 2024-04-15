package models

import (
	"gorm.io/gorm"
)

type DbBackup struct {
	Model
	FileName string `json:"file_name" valid:"required"`
	DbId     uint   `json:"-" gorm:"foreignKey:DbId" valid:"required"`
}

type DbBackupWithDbName struct {
	DbBackup
	DatabaseName string `json:"database_name"`
}

func (db *DbBackup) Create(conn *gorm.DB) error {
	return conn.Model(DbBackup{}).
		Create(&db).Error
}

func (db *DbBackup) GetById(conn *gorm.DB) (*DbBackupWithDbName, error) {
	var dbBackupWithDbName DbBackupWithDbName
	if err := conn.Table("db_backups").
		Select("db_backups.*, dbs.database_name").
		Joins("JOIN dbs ON dbs.id = db_backups.db_id").
		Where("db_backups.id = ?", db.ID).
		First(&dbBackupWithDbName).Error; err != nil {
		return nil, err
	}

	return &dbBackupWithDbName, nil
}
