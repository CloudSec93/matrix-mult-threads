package matrix

func Multiply(A, B [][]int) [][]int {
    r1 := len(A)
    c1 := len(A[0])
    c2 := len(B[0])

    result := make([][]int, r1)
    for i := range result {
        result[i] = make([]int, c2)
    }

    for i := 0; i < r1; i++ {
        for j := 0; j < c2; j++ {
            for k := 0; k < c1; k++ {
                result[i][j] += A[i][k] * B[k][j]
            }
        }
    }
    return result
}
