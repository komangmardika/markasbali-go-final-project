package services

import (
	"log"
	"time"
)

func logDuration(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func RestoreDbNoChannel() error {
	start := time.Now()
	defer logDuration(start, "RestoreDbNoChannel")
	dbs, err := ReadDatabasesJson()
	if err != nil {
		_ = SendErrorToWebSocketServer(err.Error())
		return err
	}

	configs, err := RequestLatestBackupInfo(dbs)
	if err != nil {
		_ = SendErrorToWebSocketServer(err.Error())
		return err
	}

	for _, config := range configs {
		conn, err := DownloadFile(config)
		if err != nil {
			_ = SendErrorToWebSocketServer(err.Error())
			return err
		}
		conn, err = UnzipFile(conn)
		if err != nil {
			_ = SendErrorToWebSocketServer(err.Error())
			return err
		}
		err = ImportMySQLDump(conn)
		if err != nil {
			_ = SendErrorToWebSocketServer(err.Error())
			return err
		}
	}

	return nil
}
