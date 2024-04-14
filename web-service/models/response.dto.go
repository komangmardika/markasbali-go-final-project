package models

import "time"

type DataListDto struct {
	DatabaseName string       `json:"database_name"`
	LatestBackup LatestBackup `json:"latest_backup"`
}

type DataSingleDto struct {
	DatabaseName string         `json:"database_name"`
	Histories    []LatestBackup `json:"histories"`
}

type LatestBackup struct {
	ID        uint      `json:"id"`
	FileName  string    `json:"file_name"`
	Timestamp time.Time `json:"timestamp"`
}
