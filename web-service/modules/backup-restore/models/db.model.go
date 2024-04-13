package models

import (
	"errors"
	"gorm.io/gorm"
	"markasbali_go_final_project/web-service/modules/bases/models"
	"markasbali_go_final_project/web-service/modules/constants"
)

type Db struct {
	models.Model
	DatabaseName string     `json:"database_name" valid:"required"`
	DbBackup     []DbBackup `json:"db_backups"`
}

func (db *Db) Create(conn *gorm.DB) error {
	return conn.Model(Db{}).
		Create(&db).Error
}

func (db *Db) GetById(conn *gorm.DB) (Db, error) {
	return Db{}, conn.Model(Db{}).
		Where("id = ?", db.ID).
		Where("deleted_at IS NOT NULL").
		Take(&db).Error
}

func (db *Db) GetAll(conn *gorm.DB) ([]Db, error) {
	var dbs []Db
	return dbs, conn.Find(&dbs).Error
}

func (db *Db) GetOneDatabaseWithAllBackup(conn *gorm.DB) error {
	return conn.Preload("DbBackup").Where("database_name = ?", db.DatabaseName).Take(&db).Error
}

func (db *Db) GetAllWithLatestBackup(conn *gorm.DB) ([]Db, error) {
	var dbs []Db

	subQuery := conn.Model(&DbBackup{}).
		Select("db_id, MAX(updated_at) AS latest_updated_at").
		Group("db_id")

	if err := conn.
		Preload("DbBackup", func(db *gorm.DB) *gorm.DB {
			return db.Joins("INNER JOIN (?) AS "+
				"latest_db_backups ON db_backups.db_id = latest_db_backups.db_id "+
				"AND db_backups.updated_at = latest_db_backups.latest_updated_at",
				subQuery)
		}).
		Find(&dbs).Error; err != nil {
		return nil, err
	}

	return dbs, nil
}

func (db *Db) GetAllPaginated(conn *gorm.DB, page int) ([]Db, error) {
	var dbs []Db
	pageSize := constants.DB_PAGE_SIZE

	offset := (page - 1) * pageSize

	return dbs, conn.Limit(pageSize).
		Offset(offset).
		Find(&dbs).Error
}

func (db *Db) Update(conn *gorm.DB) error {
	return conn.Model(Db{}).Select("database_name", "updated_at").
		Where("id = ?", db.Model.ID).
		Updates(map[string]interface{}{
			"database_name": db.DatabaseName,
			"updated_at":    db.UpdatedAt,
		}).Error
}

func (db *Db) Upsert(conn *gorm.DB) error {

	result := conn.Model(Db{}).Where("database_name = ?", db.DatabaseName).Take(&db)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result := conn.Create(db)
		if result.Error != nil {
			return result.Error
		}
	} else {
		result := conn.Model(&db).Updates(db)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
