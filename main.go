package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	u "workspace/users"
)

var db *sql.DB

type Env struct {
	db *sql.DB
}

// LOAD ENVIRONMENT FILE AND SET CONNECTION TO USERS PACKAGE
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectDB()
}

func injectDB(db *sql.DB, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "db", db)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// MAIN FUNCION
func main() {
	handleRequests()
}

// HANDLE REQUESTS
func handleRequests() {
	http.Handle("/users", injectDB(db, u.AllUsers))
	http.Handle("/users/create", injectDB(db, u.CreateUser))
	http.Handle("/users/update", injectDB(db, u.UpdateUser))
	http.Handle("/users/byid", injectDB(db, u.UserById))
	http.Handle("/users/delete", injectDB(db, u.DeleteUser))

	err := http.ListenAndServe(":9090", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

// CONNECT DATABASE
func connectDB() {
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USERNAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST"),
		DBName:               os.Getenv("DB_DATABASE"),
		AllowNativePasswords: true,
	}

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
