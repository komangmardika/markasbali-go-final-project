package services

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"markasbali_go_final_project/cli-service/modules/common/models"
)

func ResetDb(DbConn models.MySqlConn) error {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		DbConn.DbUsername,
		DbConn.DbPassword,
		DbConn.DbHost,
		DbConn.DbPort,
		DbConn.DatabaseName,
	)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return err
	}

	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	if err != nil {
		return err
	}

	// Drop tables in the database
	_, err = db.Exec("SHOW TABLES")
	if err != nil {
		return err
	}

	var tableName string
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		if err := rows.Scan(&tableName); err != nil {
			return err
		}
		_, err = db.Exec(fmt.Sprintf("DROP TABLE %s", tableName))
		if err != nil {
			return err
		}
	}

	// Enable foreign key checks
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}
