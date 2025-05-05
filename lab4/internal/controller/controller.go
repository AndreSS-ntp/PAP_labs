package controller

import "github.com/AndreSS-ntp/PAP_labs/tree/main/lab4/internal/model"

func ProcessMatrix(matrix model.Matrix) model.MatrixReply {
	size := len(matrix.Rows)
	reply := model.MatrixReply{}

	reply.Original = copyMatrix(matrix)

	mainDiagMin := matrix.Rows[0].Cols[0]
	secondaryDiagMin := matrix.Rows[0].Cols[size-1]

	for i := 0; i < size; i++ {
		if matrix.Rows[i].Cols[i] < mainDiagMin {
			mainDiagMin = matrix.Rows[i].Cols[i]
		}
		if matrix.Rows[i].Cols[size-1-i] < secondaryDiagMin {
			secondaryDiagMin = matrix.Rows[i].Cols[size-1-i]
		}
	}

	var targetDiag int
	if mainDiagMin < secondaryDiagMin {
		targetDiag = 0 // главная диагональ
		reply.MinValue = mainDiagMin
	} else {
		targetDiag = 1 // побочная диагональ
		reply.MinValue = secondaryDiagMin
	}

	reply.Processed = copyMatrix(matrix)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if targetDiag == 0 && i == j {
				// Главная диагональ - заменяем нулями
				reply.Processed.Rows[i].Cols[j] = 0
			} else if targetDiag == 1 && j == size-1-i {
				// Побочная диагональ - заменяем нулями
				reply.Processed.Rows[i].Cols[j] = 0
			} else if (targetDiag == 0 && i > j) || (targetDiag == 1 && i+j > size-1) {
				// Элементы ниже диагонали - в квадрат
				val := reply.Processed.Rows[i].Cols[j]
				reply.Processed.Rows[i].Cols[j] = val * val
			}
		}
	}

	return reply
}

func copyMatrix(src model.Matrix) model.Matrix {
	size := len(src.Rows)
	dst := model.Matrix{Rows: make([]model.Row, size)}
	for i := range dst.Rows {
		dst.Rows[i].Cols = make([]int, size)
		copy(dst.Rows[i].Cols, src.Rows[i].Cols)
	}
	return dst
}
