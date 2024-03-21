package main

import (
	"context"
	"devcode-todolist-api/internal/infrastructures/http"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
)

func main() {
	// read config
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln(err)
	}

	port := os.Getenv("APP_PORT")

	// listener for server
	listener, err := net.Listen(
		"tcp", fmt.Sprintf(":%s", port),
	)
	if err != nil {
		log.Fatalln(err)
	}

	// run server
	app, err := http.CreateServer(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	err = app.Serve(listener)

	if err != nil {
		log.Fatalln(err)
	}
}
