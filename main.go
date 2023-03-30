package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	connectDB()
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/users", AllUsers)
	http.HandleFunc("/users/create", CreateUser)
	http.HandleFunc("/users/update", UpdateUser)
	http.HandleFunc("/users/byid", UserById)
	http.HandleFunc("/users/delete", DeleteUser)

	err := http.ListenAndServe(":9090", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func connectDB() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USERNAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}
