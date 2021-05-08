package main

import (
	"log"
)

const addr = "0.0.0.0:8080"

func main() {
	// you could insert your favorite logger here for structured or leveled logging
	log.Println("Starting server...")

	// initialize data sources
	ds, err := initDS()
	if err != nil {
		log.Fatalf("Unable to initialize data sources: %v\n", err)
	}

	server, err := inject(ds)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	err = server.Start(addr)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
