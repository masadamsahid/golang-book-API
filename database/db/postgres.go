package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DBconn *sql.DB
	err    error
)

func ConnectPg() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

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
		fmt.Println("Failed to establish connection to DB")
		panic(err)
	}

	err = DBconn.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success connecting to DB")
}

func StopDBConn() {
	DBconn.Close()
	fmt.Println("Success closing connection to DB")
}
