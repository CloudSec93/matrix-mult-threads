package main

import (
    "fmt"
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

    fmt.Println("Multiplying matrices...")
    result := matrix.Multiply(A, B)

    fmt.Println("Result:")
    matrix.PrintMatrix(result)
}
