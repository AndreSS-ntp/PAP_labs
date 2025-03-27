package matrixops

// MatrixArgs represents the input matrix sent from client to server
type MatrixArgs struct {
	Matrix [][]int
}

// MatrixResult represents the result returned from server to client
type MatrixResult struct {
	OriginalMatrix [][]int
	ResultMatrix   [][]int
	MinDiagElement int
	MinDiagIndex   int // 0 for main diagonal, 1 for secondary diagonal
}

// MatrixService defines the service for matrix operations
type MatrixService struct{}

// ProcessMatrix is the RPC method that processes the matrix according to requirements
func (m *MatrixService) ProcessMatrix(args MatrixArgs, reply *MatrixResult) error {
	// Copy the original matrix to the result
	reply.OriginalMatrix = make([][]int, len(args.Matrix))
	for i := range args.Matrix {
		reply.OriginalMatrix[i] = make([]int, len(args.Matrix[i]))
		copy(reply.OriginalMatrix[i], args.Matrix[i])
	}

	// Create a copy of the matrix to work with
	n := len(args.Matrix)
	matrix := make([][]int, n)
	for i := range args.Matrix {
		matrix[i] = make([]int, n)
		copy(matrix[i], args.Matrix[i])
	}

	// Find minimum element in both diagonals
	mainDiagMin := matrix[0][0]
	secondaryDiagMin := matrix[0][n-1]

	// Check main diagonal
	for i := 0; i < n; i++ {
		if matrix[i][i] < mainDiagMin {
			mainDiagMin = matrix[i][i]
		}
	}

	// Check secondary diagonal
	for i := 0; i < n; i++ {
		if matrix[i][n-1-i] < secondaryDiagMin {
			secondaryDiagMin = matrix[i][n-1-i]
		}
	}

	// Determine which diagonal has the minimum element
	minDiagIndex := 0
	minDiagElement := mainDiagMin
	if secondaryDiagMin < mainDiagMin {
		minDiagIndex = 1
		minDiagElement = secondaryDiagMin
	}

	reply.MinDiagElement = minDiagElement
	reply.MinDiagIndex = minDiagIndex

	// Create the result matrix (starting with a copy of the original)
	reply.ResultMatrix = make([][]int, n)
	for i := range matrix {
		reply.ResultMatrix[i] = make([]int, n)
		copy(reply.ResultMatrix[i], matrix[i])
	}

	// Replace elements in the diagonal with minimum element with zeros
	if minDiagIndex == 0 {
		// Main diagonal
		for i := 0; i < n; i++ {
			reply.ResultMatrix[i][i] = 0
		}

		// Square elements below the main diagonal
		for i := 1; i < n; i++ {
			for j := 0; j < i; j++ {
				reply.ResultMatrix[i][j] = reply.ResultMatrix[i][j] * reply.ResultMatrix[i][j]
			}
		}
	} else {
		// Secondary diagonal
		for i := 0; i < n; i++ {
			reply.ResultMatrix[i][n-1-i] = 0
		}

		// Square elements below the secondary diagonal
		for i := 1; i < n; i++ {
			for j := n - i; j < n; j++ {
				reply.ResultMatrix[i][j] = reply.ResultMatrix[i][j] * reply.ResultMatrix[i][j]
			}
		}
	}

	return nil
}

// ProcessMatrixAndPrint is an RPC method that processes the matrix and returns results for printing
func (m *MatrixService) ProcessMatrixAndPrint(args MatrixArgs, reply *MatrixResult) error {
	// Call the original service method
	err := m.ProcessMatrix(args, reply)
	return err
}

// PrintMatrix is a helper function to print a matrix
func PrintMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			print(val, " ")
		}
		println()
	}
	println()
}
