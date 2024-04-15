package models

type MySqlConn struct {
	DatabaseName string `json:"database_name"`
	DbHost       string `json:"db_host"`
	DbPort       string `json:"db_port"`
	DbUsername   string `json:"db_username"`
	DbPassword   string `json:"db_password"`
}

type MySqlConnWithBackup struct {
	DatabaseName string `json:"database_name"`
	DbHost       string `json:"db_host"`
	DbPort       string `json:"db_port"`
	DbUsername   string `json:"db_username"`
	DbPassword   string `json:"db_password"`
	FileId       uint   `json:"file_id"`
	FileName     string `json:"file_name"`
	SqlFileName  string `json:"sql_file_name"`
	TmpFolder    string `json:"tmp_folder"`
}
