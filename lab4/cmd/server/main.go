package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab4/pkg/matrixops"
	"github.com/divan/gorilla-xmlrpc/xml"
	"github.com/fatih/color"
	"github.com/gorilla/rpc"
)

// Helper function to print a matrix
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println()
}

func main() {
	// Create a new RPC server
	s := rpc.NewServer()

	// Register XML codec
	s.RegisterCodec(xml.NewCodec(), "text/xml")

	// Create a new instance of the matrix service
	matrixService := new(matrixops.MatrixService)

	// Register the service
	err := s.RegisterService(matrixService, "")
	if err != nil {
		log.Fatal("Error registering matrix service:", err)
	}

	// Create HTTP handler for RPC
	http.Handle("/RPC2", s)

	// Print server information
	color.Green("Matrix processing server started on port 1234")
	color.Yellow("Waiting for client connections...")

	// Start serving HTTP requests
	err = http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal("Error serving:", err)
	}
}
