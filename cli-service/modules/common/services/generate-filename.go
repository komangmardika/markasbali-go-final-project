package services

import (
	"github.com/google/uuid"
	"time"
)

func GenerateFileName(dbName string) string {
	uid := uuid.New()
	currentTime := time.Now()
	formattedDateTime := currentTime.Format("2006-01-02-15-04-05")
	return "mysql-" + formattedDateTime + "-" + dbName + "-" + uid.String()
}
