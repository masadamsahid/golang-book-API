package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	DBconn *sql.DB
	err    error
)

func ConnectPg() {

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=%s`,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
	)

	DBconn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Failed to establish connection to DB", err)
	}

	err = DBconn.Ping()
	if err != nil {
		log.Println(err)
	}
	log.Println("Success connecting to DB")
}

func StopDBConn() {
	DBconn.Close()
	log.Println("Success closing connection to DB")
}
