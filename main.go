package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"github.com/raynine/ticket-system-backend/server"
)

func main() {
	log.Println("-- Loading Env variables -- ")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("-- Creating Server -- ")

	router := server.CreateServer()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	loggingMiddleware := handlers.LoggingHandler(os.Stdout, router)

	log.Printf("Server running on port: %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), loggingMiddleware)

}
