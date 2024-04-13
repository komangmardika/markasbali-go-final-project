package services

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"markasbali_go_final_project/web-service/configs"
	"markasbali_go_final_project/web-service/modules/backup-restore/models"
	"mime/multipart"
)

func BackupSqlFile(ctx *fiber.Ctx, file *multipart.FileHeader, dbName string) (models.Db, models.DbBackup, error) {

	// check create folder if it doesn't exist with db name
	path, err := CheckOrCreateFolder(dbName)

	if err != nil {
		return models.Db{}, models.DbBackup{}, err
	}

	// copy file to folder
	err = ctx.SaveFile(file, path+"/"+file.Filename)
	if err != nil {
		return models.Db{}, models.DbBackup{}, err
	}

	// sql upsert master

	db := models.Db{DatabaseName: dbName}

	isValid, err := govalidator.ValidateStruct(db)
	if !isValid {
		return models.Db{}, models.DbBackup{}, err
	}

	err = db.Upsert(configs.Mysql.DB)

	// sql insert details
	dbBackup := models.DbBackup{
		FileName: file.Filename,
		DbId:     db.ID,
	}

	isValid, err = govalidator.ValidateStruct(dbBackup)
	if !isValid {
		return models.Db{}, models.DbBackup{}, err
	}
	err = dbBackup.Create(configs.Mysql.DB)
	if err != nil {
		return models.Db{}, models.DbBackup{}, err
	}

	return db, dbBackup, nil
}
