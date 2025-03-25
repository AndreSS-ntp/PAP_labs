package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
	"strings"

	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab4/pkg/matrixops"
	"github.com/fatih/color"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer client.Close()

	color.Green("Connected to matrix processing server")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter matrix size (n for n×n matrix): ")
	sizeStr, _ := reader.ReadString('\n')
	sizeStr = strings.TrimSpace(sizeStr)

	size, err := strconv.Atoi(sizeStr)
	if err != nil || size <= 0 {
		log.Fatal("Invalid matrix size. Please enter a positive integer.")
	}

	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	fmt.Printf("Enter the elements of the %dx%d matrix row by row:\n", size, size)

	for i := 0; i < size; i++ {
		fmt.Printf("Row %d (example: 1 2 3): ", i+1)
		rowStr, _ := reader.ReadString('\n')
		rowStr = strings.TrimSpace(rowStr)

		elements := strings.Fields(rowStr)

		if len(elements) != size {
			log.Fatalf("Expected %d elements for row %d, but got %d", size, i+1, len(elements))
		}

		for j, element := range elements {
			val, err := strconv.Atoi(element)
			if err != nil {
				log.Fatalf("Invalid integer at row %d, column %d: %s", i+1, j+1, element)
			}
			matrix[i][j] = val
		}
	}

	args := &matrixops.MatrixArgs{Matrix: matrix}
	var result matrixops.MatrixResult

	err = client.Call("MatrixServiceWrapper.ProcessMatrixAndPrint", args, &result)
	if err != nil {
		log.Fatal("Error calling RPC:", err)
	}

	fmt.Println("\nResults from server:")

	color.Cyan("Original Matrix:")
	printMatrix(result.OriginalMatrix)

	color.Yellow("Minimum diagonal element: %d (on %s diagonal)",
		result.MinDiagElement,
		[]string{"main", "secondary"}[result.MinDiagIndex])

	color.Cyan("Processed Matrix:")
	printMatrix(result.ResultMatrix)
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%4d ", val)
		}
		fmt.Println()
	}
	fmt.Println()
}
