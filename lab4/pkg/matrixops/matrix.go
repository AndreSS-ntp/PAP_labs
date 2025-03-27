package matrixops

type MatrixArgs struct {
	Matrix [][]int
}

type MatrixResult struct {
	OriginalMatrix [][]int
	ResultMatrix   [][]int
	MinDiagElement int
	MinDiagIndex   int // 0 for main diagonal, 1 for secondary diagonal
}

type MatrixService struct{}

func (m *MatrixService) ProcessMatrix(args *MatrixArgs, result *MatrixResult) error {
	result.OriginalMatrix = make([][]int, len(args.Matrix))
	for i := range args.Matrix {
		result.OriginalMatrix[i] = make([]int, len(args.Matrix[i]))
		copy(result.OriginalMatrix[i], args.Matrix[i])
	}

	n := len(args.Matrix)
	matrix := make([][]int, n)
	for i := range args.Matrix {
		matrix[i] = make([]int, n)
		copy(matrix[i], args.Matrix[i])
	}

	// Find minimum
	mainDiagMin := matrix[0][0]
	secondaryDiagMin := matrix[0][n-1]

	for i := 0; i < n; i++ {
		if matrix[i][i] < mainDiagMin {
			mainDiagMin = matrix[i][i]
		}
	}

	for i := 0; i < n; i++ {
		if matrix[i][n-1-i] < secondaryDiagMin {
			secondaryDiagMin = matrix[i][n-1-i]
		}
	}

	minDiagIndex := 0
	minDiagElement := mainDiagMin
	if secondaryDiagMin < mainDiagMin {
		minDiagIndex = 1
		minDiagElement = secondaryDiagMin
	}

	result.MinDiagElement = minDiagElement
	result.MinDiagIndex = minDiagIndex

	result.ResultMatrix = make([][]int, n)
	for i := range matrix {
		result.ResultMatrix[i] = make([]int, n)
		copy(result.ResultMatrix[i], matrix[i])
	}

	// Replace elements in the diagonal with minimum element with zeros
	if minDiagIndex == 0 {
		// Main diagonal
		for i := 0; i < n; i++ {
			result.ResultMatrix[i][i] = 0
		}

		// Square elements below the main diagonal
		for i := 1; i < n; i++ {
			for j := 0; j < i; j++ {
				result.ResultMatrix[i][j] = result.ResultMatrix[i][j] * result.ResultMatrix[i][j]
			}
		}
	} else {
		// Secondary diagonal
		for i := 0; i < n; i++ {
			result.ResultMatrix[i][n-1-i] = 0
		}

		// Square elements below the secondary diagonal
		for i := 1; i < n; i++ {
			for j := n - i; j < n; j++ {
				result.ResultMatrix[i][j] = result.ResultMatrix[i][j] * result.ResultMatrix[i][j]
			}
		}
	}

	return nil
}

func PrintMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			print(val, " ")
		}
		println()
	}
	println()
}
