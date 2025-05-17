package matrix

import "fmt"

func ReadMatrix(rows, cols int) [][]int {

	// ReadMatrix does 3 things. 
	// 1. It creates a matrix of size rows x cols (we do not create the matrix in the main function)
	// 2. It fills the matrix with user input
	// 3. It returns the matrix

	
    matrix := make([][]int, rows)
    for i := range matrix {
        matrix[i] = make([]int, cols)
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            fmt.Printf("Enter element [%d][%d]: ", i, j)
            fmt.Scan(&matrix[i][j])
        }
    }
    return matrix
}

func PrintMatrix(m [][]int) {
    for i := 0; i < len(m); i++ {
        for j := 0; j < len(m[0]); j++ {
            fmt.Print(m[i][j], " ")
        }
        fmt.Println()
    }
}
