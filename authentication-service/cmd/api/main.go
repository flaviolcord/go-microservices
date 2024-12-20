package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/flaviolcord/go-microservices/auth-service/data"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	webPort    = "80"
	maxRetries = 10
)

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")

	// connect to DB
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	// set up config
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for retries := 0; retries < maxRetries; retries++ {
		connection, err := openDB(dsn)
		if err == nil {
			log.Println("Connected to Postgres!")
			return connection
    }

    log.Printf("Postgres not ready, attempt %d/%d: %v", retries+1, maxRetries, err)
    time.Sleep(time.Duration(retries+1) * time.Second)
  }

  log.Printf("failed to connect to Postgres after %d attempts", maxRetries)
  return nil 
}
