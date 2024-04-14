package services

import (
	"encoding/csv"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"markasbali_go_final_project/cli-service/models"
	"os"
	"strconv"
	"sync"
)

type MysqlDB struct {
	DB *gorm.DB
}

var Mysql MysqlDB

func ImportBook() error {

	file, err := os.Open("modules/reset/seeders/books.csv")

	if err != nil {
		return err
	}

	rowChannel := make(chan []string)
	bookChannel := make(chan models.Book)
	done := make(chan bool)
	var wg sync.WaitGroup

	go func() {
		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			panic(err)
		}

		for _, record := range records {
			rowChannel <- record
		}
		close(rowChannel)
	}()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for row := range rowChannel {
				book, err := parseRowToBook(row)
				if err == nil {
					bookChannel <- book
				}

			}
		}()
	}

	go func() {
		wg.Wait()
		close(bookChannel)
		done <- true
	}()

	// Process book yang masuk ke dalam channel book untuk disimpan di slice books
	for book := range bookChannel {
		if err = book.Upsert(Mysql.DB); err != nil {
			return err
		}
	}

	<-done

	err = file.Close()

	return err

}

func ImportCar() error {

	file, err := os.Open("modules/reset/seeders/cars_500.csv")

	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	rowChannel := make(chan []string)
	carChannel := make(chan models.Car)
	done := make(chan bool)
	var wg sync.WaitGroup

	go func() {
		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			panic(err)
		}

		for _, record := range records {
			rowChannel <- record
		}
		close(rowChannel)
	}()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for row := range rowChannel {
				car, err := parseRowToCar(row)
				if err == nil {
					carChannel <- car
				}

			}
		}()
	}

	go func() {
		wg.Wait()
		close(carChannel)
		done <- true
	}()

	// Process book yang masuk ke dalam channel book untuk disimpan di slice books
	for car := range carChannel {
		if err = car.Upsert(Mysql.DB); err != nil {
			fmt.Println(err)
		}
	}

	<-done

	return err

}

func parseRowToBook(row []string) (models.Book, error) {

	id, err := strconv.ParseUint(row[0], 10, 64)
	if err != nil {
		return models.Book{}, err
	}
	year, _ := strconv.ParseUint(row[3], 10, 64)
	stock, _ := strconv.ParseUint(row[6], 10, 64)

	return models.Book{
		Model:   models.Model{ID: uint(id)},
		ISBN:    row[1],
		Penulis: row[2],
		Tahun:   uint(year),
		Judul:   row[4],
		Gambar:  row[5],
		Stok:    uint(stock),
	}, nil

}

func parseRowToCar(row []string) (models.Car, error) {
	// Convert Tahun and Stok to uint
	year, err := strconv.ParseUint(row[1], 10, 64)

	if err != nil {
		return models.Car{}, err
	}
	return models.Car{
		Name:    row[2],
		CarType: row[3],
		Year:    uint(year),
		Uuid:    row[0],
	}, nil

}

func OpenDB(silentLogger bool, conn models.MySqlConn) {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		conn.DbUsername,
		conn.DbPassword,
		conn.DbHost,
		conn.DbPort,
		conn.DatabaseName,
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

}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Book{}, &models.Car{})

	if err != nil {
		return err
	}

	return err
}

func CloseDb(db *gorm.DB) error {
	dbx, _ := db.DB()
	err := dbx.Close()
	if err != nil {
		return err
	}

	return nil
}
