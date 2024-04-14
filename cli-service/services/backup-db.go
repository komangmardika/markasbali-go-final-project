package services

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"markasbali_go_final_project/cli-service/models"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
)

func BackupDb() error {
	configs, _ := ReadDatabasesJson()
	numWorkers := 2

	jobCh := make(chan models.MySqlConn, len(configs))

	for i := 0; i < numWorkers; i++ {
		go func() {
			err := doBackup(jobCh)
			if err != nil {

			}
		}()
	}

	for _, db := range configs {
		jobCh <- db
	}
	close(jobCh)
	return nil
}

func runMySQLDump(db models.MySqlConn, fileName string) (string, error) {

	sqlFileName := os.Getenv("TMP_FOLDER_PATH") + fileName + ".sql"
	cmd := exec.Command("mysqldump", "-u", db.DbUsername, "-p"+db.DbPassword, db.DatabaseName)

	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	// Write SQL data to file
	err = os.WriteFile(sqlFileName, out, 0777)
	if err != nil {
		return "", err
	}

	return sqlFileName, nil
}

func zipSqlFile(sqlFileName string, fileName string) error {

	zipFileName := os.Getenv("TMP_FOLDER_PATH") + fileName + ".zip"
	zipFile, err := os.Create(zipFileName)

	if err != nil {
		return err
	}

	zipWriter := zip.NewWriter(zipFile)

	// Add SQL file to zip archive
	sqlFile, err := os.Open(sqlFileName)
	if err != nil {
		return err
	}

	fileInfo, err := sqlFile.Stat()
	if err != nil {
		err := sqlFile.Close()
		if err != nil {
			return err
		}
	}

	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {

		err := sqlFile.Close()
		if err != nil {
			return err
		}
	}

	header.Name = sqlFileName
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		err := sqlFile.Close()
		if err != nil {
			return err
		}
	}

	_, err = io.Copy(writer, sqlFile)
	if err != nil {
		err := sqlFile.Close()
		if err != nil {
			return err
		}
	}

	err = sqlFile.Close()
	if err != nil {
		return err
	}
	err = zipWriter.Close()
	if err != nil {
		return err
	}
	err = zipFile.Close()
	if err != nil {
		return err
	}

	return nil
}

func SendFileToWebService(zipName string, dbName string) error {
	body := &bytes.Buffer{}

	file, err := os.Open(os.Getenv("TMP_FOLDER_PATH") + zipName + ".zip")
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// Create a new multipart writer
	writer := multipart.NewWriter(body)

	// Create a new form file field
	fileWriter, err := writer.CreateFormFile("file", os.Getenv("TMP_FOLDER_PATH")+zipName+".zip")
	if err != nil {
		return err
	}

	// Copy the file content to the form file field
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return err
	}

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		return err
	}

	// Create a new HTTP request
	webServiceBaseUrl := fmt.Sprintf("%s:%s/", os.Getenv("APP_WEB_SERVICE_URL"), os.Getenv("APP_WEB_SERVICE_PORT"))
	req, err := http.NewRequest("POST", webServiceBaseUrl+"web/"+dbName, body)
	if err != nil {
		return err
	}

	// Set the content type header
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "X-TOKEN "+base64.StdEncoding.EncodeToString([]byte(os.Getenv("SECRET_TOKEN"))))

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		err = os.Remove(os.Getenv("TMP_FOLDER_PATH") + zipName + ".zip")
		if err != nil {
			return
		}
		err = os.Remove(os.Getenv("TMP_FOLDER_PATH") + zipName + ".sql")
		if err != nil {
			return
		}
	}(resp.Body)

	return nil
}

func doBackup(jobCh <-chan models.MySqlConn) error {
	for db := range jobCh {

		// Dump MySQL database to SQL file
		fileName := GenerateFileName(db.DatabaseName)

		sqlFileName, err := runMySQLDump(db, fileName)
		if err != nil {
			return err
		}

		// Create a zip archive
		err = zipSqlFile(sqlFileName, fileName)
		if err != nil {
			return err
		}

		err = SendFileToWebService(fileName, db.DatabaseName)

		if err != nil {
			return err
		}
	}

	return nil
}
