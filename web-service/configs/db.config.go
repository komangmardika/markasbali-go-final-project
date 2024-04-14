package configs

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"markasbali_go_final_project/web-service/models"
	"os"
)

type MysqlDB struct {
	DB *gorm.DB
}

var Mysql MysqlDB

func OpenDB(silentLogger bool) {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	MysqlConn, err := gorm.Open(mysql.Open(connString), &gorm.Config{})

	if silentLogger {
		MysqlConn, err = gorm.Open(mysql.Open(connString), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
	}

	if err != nil {
		log.Fatal(err)
	}

	Mysql = MysqlDB{
		DB: MysqlConn,
	}

	err = autoMigrate(Mysql.DB)
	if err != nil {
		return
	}
}

func autoMigrate(db *gorm.DB) error {

	err := db.AutoMigrate(&models.Db{}, &models.DbBackup{})
	if err != nil {
		return err
	}
	db.Model(&models.Db{}).Association("Histories")

	if err != nil {
		return err
	}

	return err
}
