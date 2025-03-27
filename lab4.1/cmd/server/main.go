package main

import (
	"fmt"
	"log"
	"math"
	"net"
	"net/http"
	"net/rpc"
)

type MatrixService struct{}

type MatrixRequest struct {
	Matrix [][]int
}

type MatrixResponse struct {
	OriginalMatrix  [][]int
	ProcessedMatrix [][]int
	MinDiagElement  int
	MinDiagIndex    int
}

func (m *MatrixService) ProcessMatrix(req MatrixRequest, resp *MatrixResponse) error {
	// Copy the original matrix to the response
	resp.OriginalMatrix = make([][]int, len(req.Matrix))
	for i := range req.Matrix {
		resp.OriginalMatrix[i] = make([]int, len(req.Matrix[i]))
		copy(resp.OriginalMatrix[i], req.Matrix[i])
	}

	// Create a new matrix for processing
	matrix := make([][]int, len(req.Matrix))
	for i := range req.Matrix {
		matrix[i] = make([]int, len(req.Matrix[i]))
		copy(matrix[i], req.Matrix[i])
	}

	// Find the diagonal with the minimum element
	minElement := math.MaxInt32
	minDiagIndex := 0
	size := len(matrix)

	// Check main diagonal
	for i := 0; i < size; i++ {
		if matrix[i][i] < minElement {
			minElement = matrix[i][i]
			minDiagIndex = 0 // Main diagonal
		}
	}

	// Check secondary diagonal
	for i := 0; i < size; i++ {
		if matrix[i][size-1-i] < minElement {
			minElement = matrix[i][size-1-i]
			minDiagIndex = 1 // Secondary diagonal
		}
	}

	resp.MinDiagElement = minElement
	resp.MinDiagIndex = minDiagIndex

	// Process the matrix based on the diagonal with minimum element
	if minDiagIndex == 0 {
		// Main diagonal has minimum element
		// Set main diagonal elements to zero
		for i := 0; i < size; i++ {
			matrix[i][i] = 0
		}

		// Square elements below the main diagonal
		for i := 0; i < size; i++ {
			for j := 0; j < i; j++ {
				matrix[i][j] = matrix[i][j] * matrix[i][j]
			}
		}
	} else {
		// Secondary diagonal has minimum element
		// Set secondary diagonal elements to zero
		for i := 0; i < size; i++ {
			matrix[i][size-1-i] = 0
		}

		// Square elements below the secondary diagonal
		for i := 0; i < size; i++ {
			for j := size - i; j < size; j++ {
				matrix[i][j] = matrix[i][j] * matrix[i][j]
			}
		}
	}

	resp.ProcessedMatrix = matrix
	return nil
}

func main() {
	matrixService := new(MatrixService)
	rpc.Register(matrixService)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	fmt.Println("Matrix processing server started on port 1234")
	http.Serve(listener, nil)
}
