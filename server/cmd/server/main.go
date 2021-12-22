package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/zavad4/lab3go/tree/main/server/db"
)

var httpPortNumber = flag.Int("p", 8080, "HTTP port number")

func NewDbConnection() (*sql.DB, error) {
	conn := &db.Connection{
		DbName:     "d4qgrhe3its6v7",
		User:       "hcidzrejqbzxrj",
		Password:   "5c8d70b2c65db9dc3b4635dfb8140585a6a8a33582a8fdc7e36b79496d587c2d",
		Host:       "ec2-54-195-246-55.eu-west-1.compute.amazonaws.com",
		DisableSSL: false,
	}
	return conn.Open()
}

func main() {
	// Parse command line arguments. Port number may be defined with "-p" flag.
	flag.Parse()
	// Create the server.
	server := ComposeApiServer(HttpPortNumber(*httpPortNumber))
	// Start it.
	go func() {
		log.Println("Starting server...")

		err := server.Start()
		if err == http.ErrServerClosed {
			log.Printf("HTTP server stopped")
			server.Stop()
		} else {
			log.Fatalf("Cannot start HTTP server: %s", err)
		}
	}()

	// Wait for Ctrl-C signal.
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	<-sigChannel

	if err := server.Stop(); err != nil && err != http.ErrServerClosed {
		log.Printf("Error stopping the server: %s", err)
	}
}
