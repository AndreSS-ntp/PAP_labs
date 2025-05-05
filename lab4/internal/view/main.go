package view

import (
	"fmt"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab4/internal/model"
)

func PrintMatrix(matrix model.Matrix) {
	for _, row := range matrix.Rows {
		for _, val := range row.Cols {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
