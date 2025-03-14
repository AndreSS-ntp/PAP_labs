package main

import (
	"fmt"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab2/internal/server"
	"log"
)

func main() {
	serv := server.NewServer(":8080")
	fmt.Println("Starting chat server on port 8080...")
	err := serv.Start()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
