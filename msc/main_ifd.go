package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to MySQL database
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3310)/inventaris")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Insert fake data into the table
	for i := 0; i < 1000; i++ {
		var name string
		var amount int16
		var price float32

		// Generate fake name
		name = faker.Name()

		// Generate random amount
		amount = int16(rand.Intn(100) + 1) // Random number between 1 and 100

		// Generate random price between 10000 and 100000
		price = float32(rand.Intn(90001) + 10000) // Random number between 10000 and 100000

		_, err = db.Exec("INSERT INTO barang (Nama, Jumlah, Harga) VALUES (?, ?, ?)", name, amount, price)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Fake data inserted successfully")
}
