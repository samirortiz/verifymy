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

	a "workspace/auth"
	u "workspace/users"
)

var db *sql.DB

// LOAD ENVIRONMENT FILE AND SET CONNECTION TO USERS PACKAGE
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	u.SetDB(connectDB())
}

// MAIN FUNCION
func main() {
	handleRequests()
}

// HANDLE REQUESTS
func handleRequests() {
	http.HandleFunc("/users", u.AllUsers)
	http.HandleFunc("/users/create", u.CreateUser)
	http.HandleFunc("/users/update", u.UpdateUser)
	http.HandleFunc("/users/getbyid", u.UserById)
	http.HandleFunc("/users/delete", u.DeleteUser)
	http.HandleFunc("/signin", a.Signin)
	http.HandleFunc("/refresh", a.Refresh)
	http.HandleFunc("/logout", a.Logout)

	err := http.ListenAndServe(":9090", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

// CONNECT DATABASE
func connectDB() *sql.DB {
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USERNAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST"),
		DBName:               os.Getenv("DB_DATABASE"),
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open(os.Getenv("DB_DRIVER"), cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	return db
}
