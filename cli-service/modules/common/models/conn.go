package models

type MySqlConn struct {
	DatabaseName string `json:"database_name"`
	DbHost       string `json:"db_host"`
	DbPort       string `json:"db_port"`
	DbUsername   string `json:"db_username"`
	DbPassword   string `json:"db_password"`
}
