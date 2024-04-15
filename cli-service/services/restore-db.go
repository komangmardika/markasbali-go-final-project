package services

import (
	"archive/zip"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"markasbali_go_final_project/cli-service/models"
	"markasbali_go_final_project/proto"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

func RequestLatestBackupInfo(conn []models.MySqlConn) ([]models.MySqlConnWithBackup, error) {
	resp, err := RequestRestoreToServer("restore", "all-db-latest-history", conn[0].DatabaseName)
	if err != nil {
		_ = SendErrorToWebSocketServer(err.Error())
		return []models.MySqlConnWithBackup{}, err
	}
	var detail []models.DataListDto
	err = json.Unmarshal([]byte(resp), &detail)
	if err != nil {
		_ = SendErrorToWebSocketServer(err.Error())
		return []models.MySqlConnWithBackup{}, err
	}

	var connWithBackups []models.MySqlConnWithBackup

	for i := 0; i < len(conn); i++ {
		for j := 0; j < len(detail); j++ {
			if conn[i].DatabaseName == detail[j].DatabaseName {
				connWithBackups = append(connWithBackups, models.MySqlConnWithBackup{
					DatabaseName: conn[i].DatabaseName,
					DbUsername:   conn[i].DbUsername,
					DbPassword:   conn[i].DbPassword,
					DbPort:       conn[i].DbPort,
					DbHost:       conn[i].DbHost,
					FileId:       detail[j].LatestBackup.ID,
					FileName:     detail[j].LatestBackup.FileName,
					TmpFolder:    os.Getenv("TMP_FOLDER_PATH"),
					SqlFileName:  strings.Replace(detail[j].LatestBackup.FileName, "zip", "sql", -1),
				})
				break
			}
		}
	}

	return connWithBackups, nil
}

func DownloadFile(conn models.MySqlConnWithBackup) (models.MySqlConnWithBackup, error) {
	// using grpc request a file and save it to tmp folder
	resp, err := RequestFileToServer(conn.FileId)
	if err != nil {
		_ = SendErrorToWebSocketServer(err.Error())
		return conn, err
	}

	err = SaveToFile(resp, os.Getenv("TMP_FOLDER_PATH")+conn.FileName)
	if err != nil {
		_ = SendErrorToWebSocketServer(err.Error())
		return conn, err
	}

	return conn, nil
}

func UnzipFile(conn models.MySqlConnWithBackup) (models.MySqlConnWithBackup, error) {
	r, err := zip.OpenReader(conn.TmpFolder + conn.FileName)
	log.Println(conn.TmpFolder + conn.FileName)

	if err != nil {
		_ = SendErrorToWebSocketServer(err.Error())
		return conn, err
	}

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			_ = SendErrorToWebSocketServer(err.Error())
			return conn, err
		}

		path := filepath.Join("", f.Name)

		// Buat file tujuan.
		fDest, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())

		// Salin isi file dari ZIP ke file tujuan.
		_, err = io.Copy(fDest, rc)
		if err != nil {
			_ = SendErrorToWebSocketServer(err.Error())
			return conn, err
		}

		if err != nil {
			_ = SendErrorToWebSocketServer(err.Error())
			return conn, err
		}

		if err != nil {
			_ = SendErrorToWebSocketServer(err.Error())
			return conn, err
		}

		err = fDest.Close()
		if err != nil {
			_ = SendErrorToWebSocketServer(err.Error())
			return models.MySqlConnWithBackup{}, err
		}
		if err != nil {
			_ = SendErrorToWebSocketServer(err.Error())
			return conn, err
		}
	}

	err = r.Close()
	if err != nil {
		_ = SendErrorToWebSocketServer(err.Error())
		return conn, err
	}

	err = os.Remove(conn.TmpFolder + conn.FileName)
	if err != nil {
		_ = SendErrorToWebSocketServer(err.Error())
		return models.MySqlConnWithBackup{}, err
	}

	return conn, nil
}

func ImportMySQLDump(conn models.MySqlConnWithBackup) error {
	cmd := exec.Command("mysql", fmt.Sprintf("-u%s", conn.DbUsername), fmt.Sprintf("-p%s", conn.DbPassword), conn.DatabaseName)

	f := strings.Replace(conn.FileName, ".zip", ".sql", -1)
	input, err := os.Open(os.Getenv("TMP_FOLDER_PATH") + f)
	if err != nil {
		_ = SendErrorToWebSocketServer(err.Error())
		return err
	}
	cmd.Stdin = input

	// Start the command
	if err := cmd.Start(); err != nil {
		_ = SendErrorToWebSocketServer(err.Error())
		return err
	}

	// Wait for the command to finish executing
	if err := cmd.Wait(); err != nil {
		_ = SendErrorToWebSocketServer(err.Error())
		return err
	}

	// Execute the command

	err = input.Close()
	if err != nil {
		_ = SendErrorToWebSocketServer(err.Error())
		return err
	}

	if err != nil {
		_ = SendErrorToWebSocketServer(err.Error())
		return err
	}
	err = os.Remove(os.Getenv("TMP_FOLDER_PATH") + f)
	if err != nil {
		_ = SendErrorToWebSocketServer(err.Error())
		return err
	}
	return nil
}

func RestoreDb() error {
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

	downloadFileCh := make(chan models.MySqlConnWithBackup)
	unzipFileCh := make(chan models.MySqlConnWithBackup)
	restoreDumpCh := make(chan models.MySqlConnWithBackup)
	doneCh := make(chan models.MySqlConnWithBackup)

	var wg sync.WaitGroup

	// Stage 1: DownloadFile
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(downloadFileCh)

		for _, config := range configs {
			conn, err := DownloadFile(config)
			if err != nil {
				_ = SendErrorToWebSocketServer(err.Error())
				return
			}
			downloadFileCh <- conn
		}
	}()

	// Stage 2: UnzipFile
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(unzipFileCh)

		for d := range downloadFileCh {
			conn, err := UnzipFile(d)
			if err != nil {
				_ = SendErrorToWebSocketServer(err.Error())
				return
			}
			unzipFileCh <- conn
		}
	}()

	// Stage 3: RestoreMySQLDump
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(restoreDumpCh)

		for conn := range unzipFileCh {
			err := ImportMySQLDump(conn)
			if err != nil {
				_ = SendErrorToWebSocketServer(err.Error())
				return
			}
			restoreDumpCh <- conn
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(doneCh)

		for conn := range restoreDumpCh {
			doneCh <- conn
		}
	}()

	return nil
}

func Restore(ctx context.Context, client proto.RestoreServiceClient, module string, process string, dbName string) (string, error) {
	res, err := client.Restore(ctx, &proto.RestoreRequest{Module: module, Process: process, DbName: dbName})
	if err != nil {
		return "", err
	}

	return res.RestoreResponse, err
}

func File(ctx context.Context, client proto.FileServiceClient, fileId uint32) ([]byte, error) {
	res, err := client.GetFile(ctx, &proto.FileRequest{FileId: fileId})
	if err != nil {
		return []byte{}, err
	}

	return res.GetFileContent(), err
}

func RequestRestoreToServer(module string, process string, dbName string) (string, error) {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}

	restoreService := proto.NewRestoreServiceClient(conn)

	if err != nil {
		return "", err
	}

	restoreResponse, err := Restore(context.Background(), restoreService, module, process, dbName)
	err = conn.Close()
	if err != nil {
		return "", err
	}
	return restoreResponse, nil
}

func RequestFileToServer(fileId uint) ([]byte, error) {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return []byte{}, err
	}

	fileService := proto.NewFileServiceClient(conn)

	if err != nil {
		return []byte{}, err
	}

	fileResponse, err := File(context.Background(), fileService, uint32(fileId))
	err = conn.Close()
	if err != nil {
		return []byte{}, err
	}
	return fileResponse, nil
}
