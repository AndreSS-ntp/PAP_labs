package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab4/internal/model"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab4/internal/view"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run client.go <server_address>")
	}
	serverAddress := os.Args[1]

	var size int
	fmt.Print("Enter matrix size: ")
	_, err := fmt.Scan(&size)
	if err != nil {
		log.Fatal("Input error:", err)
	}

	matrix := model.Matrix{Rows: make([]model.Row, size)}
	for i := 0; i < size; i++ {
		matrix.Rows[i].Cols = make([]int, size)
		for j := 0; j < size; j++ {
			fmt.Printf("Enter element [%d][%d]: ", i, j)
			_, err := fmt.Scan(&matrix.Rows[i].Cols[j])
			if err != nil {
				log.Fatal("Input error:", err)
			}
		}
	}

	args := model.MatrixArgs{Matrix: matrix}
	xmlData, err := xml.MarshalIndent(args, "", "  ")
	if err != nil {
		log.Fatal("XML encoding error:", err)
	}

	log.Printf("Sending POST request with XML:\n%s\n", xmlData)

	resp, err := http.Post("http://"+serverAddress+"/rpc", "text/xml", bytes.NewBuffer(xmlData))
	if err != nil {
		log.Fatal("POST request error:", err)
	}
	defer resp.Body.Close()

	var reply model.MatrixReply
	err = xml.NewDecoder(resp.Body).Decode(&reply)
	if err != nil {
		log.Fatal("XML decode error:", err)
	}

	log.Println("Received GET response with results")

	fmt.Println("\nOriginal matrix:")
	view.PrintMatrix(reply.Original)
	fmt.Println("\nMin diagonal value:", reply.MinValue)
	fmt.Println("\nProcessed matrix:")
	view.PrintMatrix(reply.Processed)
}
