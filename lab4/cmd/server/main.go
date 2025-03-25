package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab4/pkg/matrixops"
	"github.com/fatih/color"
)

type MatrixServiceWrapper struct {
	Service *matrixops.MatrixService
}

// ProcessMatrixAndPrint is an RPC method that processes the matrix and prints results
func (w *MatrixServiceWrapper) ProcessMatrixAndPrint(args *matrixops.MatrixArgs, result *matrixops.MatrixResult) error {
	err := w.Service.ProcessMatrix(args, result)
	if err != nil {
		return err
	}

	fmt.Println("Server received matrix:")
	printMatrix(args.Matrix)

	fmt.Printf("Minimum diagonal element: %d (on %s diagonal)\n",
		result.MinDiagElement,
		[]string{"main", "secondary"}[result.MinDiagIndex])

	fmt.Println("Result matrix:")
	printMatrix(result.ResultMatrix)

	return nil
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println()
}

func main() {
	matrixService := new(matrixops.MatrixService)

	err := rpc.Register(matrixService)
	if err != nil {
		log.Fatal("Error registering matrix service:", err)
	}

	wrapper := &MatrixServiceWrapper{Service: matrixService}
	err = rpc.Register(wrapper)
	if err != nil {
		log.Fatal("Error registering wrapper service:", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Error listening:", err)
	}

	color.Green("Matrix processing server started on port 1234")
	color.Yellow("Waiting for client connections...")

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving:", err)
	}
}
