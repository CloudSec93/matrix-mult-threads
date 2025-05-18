package main

import (
    "fmt"
    "time"
    "github.com/CloudSec93/matrix-mult-threads/matrix"
)


func main() {
    var r1, c1, r2, c2 int

    fmt.Print("Enter rows and columns of Matrix A: ")
    fmt.Scan(&r1, &c1)
    fmt.Print("Enter rows and columns of Matrix B: ")
    fmt.Scan(&r2, &c2)

    if c1 != r2 {
        fmt.Println("Error: Column count of Matrix A must equal row count of Matrix B.")
        return
    }

    fmt.Println("Enter Matrix A:")
    A := matrix.ReadMatrix(r1, c1)

    fmt.Println("Enter Matrix B:")
    B := matrix.ReadMatrix(r2, c2)

    fmt.Println("Matrix A:")
    matrix.PrintMatrix(A)
    fmt.Println("Matrix B:")
    matrix.PrintMatrix(B)

    var choice string
    fmt.Print("Use parallel multiplication? (yes/no): ")
    fmt.Scan(&choice)

var result [][]int

if choice == "yes" {
    fmt.Println("Multiplying using goroutines...")
    start := time.Now()
    result = matrix.MultiplyParallel(A, B)
    elapsed := time.Since(start)
    fmt.Printf("Parallel multiplication took %s\n", elapsed)
} else {
    fmt.Println("Multiplying using normal method...")
    start := time.Now()
    result = matrix.Multiply(A, B)
    elapsed := time.Since(start)
    fmt.Printf("Sequential multiplication took %s\n", elapsed)
}


    fmt.Println("Result: ")
    matrix.PrintMatrix(result)
}
