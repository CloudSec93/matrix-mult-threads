package main

import "fmt"

func main() {

    // user input for dimensions of the matrix

    var r1 int
    var c1 int

    fmt.Print("Enter the number of rows in the first matrix: ")
    fmt.Scan(&r1)
    fmt.Print("Enter the number of columns in the first matrix: ")
    fmt.Scan(&c1)

    var r2 int
    var c2 int

    fmt.Print("Enter the number of rows in the second matrix: ")
    fmt.Scan(&r2)
    fmt.Print("Enter the number of columns in the second matrix: ")
    fmt.Scan(&c2)


    // edge case: check if c1 == r2

    if c1 != r2 {
        fmt.Println("Error: The number of columns in the first matrix must be equal to the number of rows in the second matrix.")
        return
    }

    // create the matrices
    // cannot do dynamic allocation because array size should be known at compile time
    
    // matrix1 := [r1][c1]int{}
    // matrix2 := [r2][c2]int{}

    // Create the matrices as slices
    matrix1 := make([][]int, r1)
    for i := range matrix1 {
        matrix1[i] = make([]int, c1)
    }  

    matrix2 := make([][]int, r2)
    for i := range matrix2 {
        matrix2[i] = make([]int, c2)
    }


    // ask the user for the elements of the first matrix

    fmt.Println("Enter the elements of the first matrix:")
    for i := 0; i < r1; i++ {
        for j := 0; j < c1; j++ {
            fmt.Println("Enter element [", i, "][", j, "]:")
            fmt.Scan(&matrix1[i][j])
        }
    }

    // ask the user for the elements of the second matrix
    fmt.Println("Enter the elements of the second matrix:")
    for i := 0; i < r2; i++ {
        for j := 0; j < c2; j++ {
            fmt.Println("Enter element [", i, "][", j, "]:")
            fmt.Scan(&matrix2[i][j])
        }
    }

    // print matrix1

    fmt.Println("Matrix 1:")
    for i := 0; i < r1; i++ {
        for j := 0; j < c1; j++ {
            fmt.Print(matrix1[i][j], " ")
        }
        fmt.Println()
    }
    // print matrix2
    fmt.Println("Matrix 2:")
    for i := 0; i < r2; i++ {
        for j := 0; j < c2; j++ {
            fmt.Print(matrix2[i][j], " ")
        }
        fmt.Println()
    }

    // multiply the matrices

    matrix3 := make([][]int, r1)
    for i := range matrix3 {
        matrix3[i] = make([]int, c2)
    }
    for i := 0; i < r1; i++ {
        for j := 0; j < c2; j++ {
            for k := 0; k < c1; k++ {
                matrix3[i][j] += matrix1[i][k] * matrix2[k][j]
            }
        }
    }

    // print the result
    fmt.Println("Resultant Matrix:")
    for i := 0; i < r1; i++ {
        for j := 0; j < c2; j++ {
            fmt.Print(matrix3[i][j], " ")
        }
        fmt.Println()
    }





    
}