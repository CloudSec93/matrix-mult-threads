package main

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/CloudSec93/matrix-mult-threads/matrix"
)

func generateRandomMatrix(rows, cols int) [][]int {
    mat := make([][]int, rows)
    for i := range mat {
        mat[i] = make([]int, cols)
        for j := range mat[i] {
            mat[i][j] = rand.Intn(10) // Random int [0-9]
        }
    }
    return mat
}

func main() {
    rows, cols := 500, 500

    fmt.Println("Generating matrices...")
    A := generateRandomMatrix(rows, cols)
    B := generateRandomMatrix(cols, rows)

    fmt.Println("Running sequential multiplication...")
    start := time.Now()
    _ = matrix.Multiply(A, B)
    durationSeq := time.Since(start)

    fmt.Println("Running parallel multiplication with goroutines...")
    start = time.Now()
    _ = matrix.MultiplyParallel(A, B)
    durationPar := time.Since(start)

    fmt.Println("\nResults:")
    fmt.Printf("Sequential: %v\n", durationSeq)
    fmt.Printf("Parallel:   %v\n", durationPar)
}
