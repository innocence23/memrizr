package main

import (
	"log"
	"memrizr/handler"
)

const addr = "0.0.0.0:8080"

func main() {
	// you could insert your favorite logger here for structured or leveled logging
	log.Println("Starting server...")

	err := handler.NewServer().Start(addr)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
