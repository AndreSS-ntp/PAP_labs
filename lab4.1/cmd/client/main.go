package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
	"strings"
)

type MatrixRequest struct {
	Matrix [][]int
}

type MatrixResponse struct {
	OriginalMatrix  [][]int
	ProcessedMatrix [][]int
	MinDiagElement  int
	MinDiagIndex    int
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%4d ", val)
		}
		fmt.Println()
	}
}

func readMatrixSize(reader *bufio.Reader) int {
	for {
		fmt.Print("Enter matrix size (n for n×n matrix): ")
		sizeStr, _ := reader.ReadString('\n')
		sizeStr = strings.TrimSpace(sizeStr)
		size, err := strconv.Atoi(sizeStr)
		if err != nil || size <= 0 {
			fmt.Println("Invalid matrix size. Please enter a positive integer.")
			continue
		}
		return size
	}
}

func readMatrixRow(reader *bufio.Reader, rowNum, size int) []int {
	row := make([]int, size)
	for {
		fmt.Printf("Row %d (space-separated integers): ", rowNum)
		rowStr, _ := reader.ReadString('\n')
		rowStr = strings.TrimSpace(rowStr)
		elements := strings.Fields(rowStr)

		if len(elements) != size {
			fmt.Printf("Error: Expected %d elements, but got %d. Please try again.\n", size, len(elements))
			continue
		}

		validRow := true
		for j, element := range elements {
			val, err := strconv.Atoi(element)
			if err != nil {
				fmt.Printf("Error: Invalid input '%s'. Please enter integers only.\n", element)
				validRow = false
				break
			}
			row[j] = val
		}

		if validRow {
			return row
		}
	}
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connection error:", err)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Matrix Processing Client")
	fmt.Println("------------------------")

	size := readMatrixSize(reader)

	matrix := make([][]int, size)

	fmt.Printf("Enter the elements of the %d×%d matrix (row by row):\n", size, size)
	for i := 0; i < size; i++ {
		matrix[i] = readMatrixRow(reader, i+1, size)
	}

	request := MatrixRequest{Matrix: matrix}
	var response MatrixResponse

	err = client.Call("MatrixService.ProcessMatrix", request, &response)
	if err != nil {
		log.Fatal("RPC error:", err)
	}

	fmt.Println("\nResults:")
	fmt.Println("--------")

	fmt.Println("Original Matrix:")
	printMatrix(response.OriginalMatrix)

	diagType := "main diagonal"
	if response.MinDiagIndex == 1 {
		diagType = "secondary diagonal"
	}

	fmt.Printf("\nMinimum element: %d (found in the %s)\n\n", response.MinDiagElement, diagType)

	fmt.Println("Processed Matrix:")
	printMatrix(response.ProcessedMatrix)
}
