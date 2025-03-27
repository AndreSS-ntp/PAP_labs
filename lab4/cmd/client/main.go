package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab4/pkg/matrixops"
	"github.com/divan/gorilla-xmlrpc/xml"
	"github.com/fatih/color"
)

func main() {
	// Create XML-RPC client
	client := new(http.Client)

	color.Green("Connected to matrix processing server")
	fmt.Println()

	// Get matrix size from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter matrix size (n for n×n matrix): ")
	sizeStr, _ := reader.ReadString('\n')
	sizeStr = strings.TrimSpace(sizeStr)

	size, err := strconv.Atoi(sizeStr)
	if err != nil || size <= 0 {
		log.Fatal("Invalid matrix size. Please enter a positive integer.")
	}

	// Initialize the matrix
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	// Get matrix elements from user
	fmt.Printf("Enter the elements of the %d×%d matrix row by row:\n", size, size)

	for i := 0; i < size; i++ {
		fmt.Printf("Row %d (space-separated integers): ", i+1)
		rowStr, _ := reader.ReadString('\n')
		rowStr = strings.TrimSpace(rowStr)

		// Split the row string into individual numbers
		elements := strings.Fields(rowStr)

		if len(elements) != size {
			log.Fatalf("Expected %d elements for row %d, but got %d", size, i+1, len(elements))
		}

		// Convert string elements to integers
		for j, element := range elements {
			val, err := strconv.Atoi(element)
			if err != nil {
				log.Fatalf("Invalid integer at row %d, column %d: %s", i+1, j+1, element)
			}
			matrix[i][j] = val
		}
	}

	// Prepare the arguments for the RPC call
	args := matrixops.MatrixArgs{Matrix: matrix}
	var result matrixops.MatrixResult

	// Create XML-RPC request
	message, err := xml.EncodeClientRequest("MatrixService.ProcessMatrixAndPrint", args)
	if err != nil {
		log.Fatal("Error encoding client request:", err)
	}

	// Send request to server
	req, err := http.NewRequest("POST", "http://localhost:1234/RPC2", bytes.NewBuffer(message))
	if err != nil {
		log.Fatal("Error creating request:", err)
	}
	req.Header.Set("Content-Type", "text/xml")

	// Get response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	// Decode response
	err = xml.DecodeClientResponse(resp.Body, &result)
	if err != nil {
		log.Fatal("Error decoding response:", err)
	}

	// Display the results
	fmt.Println("\nResults from server:")

	color.Cyan("Original Matrix:")
	printMatrix(result.OriginalMatrix)

	color.Yellow("Minimum diagonal element: %d (on %s diagonal)",
		result.MinDiagElement,
		[]string{"main", "secondary"}[result.MinDiagIndex])

	color.Cyan("Processed Matrix:")
	printMatrix(result.ResultMatrix)
}

// Helper function to print a matrix
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%4d ", val)
		}
		fmt.Println()
	}
	fmt.Println()
}
