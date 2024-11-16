package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/flaviolcord/go-microservices/auth-service/data"
)

const WebPort = "80"

type Config struct {
  DB *sql.DB
  Models data.Models
}

func main() {
  log.Println("Starting the authentication service")

  app := Config {

  }

  srv := http.Server{
    Addr: fmt.Sprintf(":%s", WebPort),
    Handler: app.routes(),
  }

  err := srv.ListenAndServe()
  if err != nil {
    log.Panic(err)
  }
}
