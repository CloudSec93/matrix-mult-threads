package matrix

import (
	"sync"
)

func MultiplyParallel(A, B [][]int) [][]int {
	r1 := len(A)
	c1 := len(A[0])
	c2 := len(B[0])

	result := make([][]int, r1)
	for i := range result {
		result[i] = make([]int, c2)
	}

	var wg sync.WaitGroup

	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			wg.Add(1) 

			// goroutine for each element of the result matrix
			go func(i, j int) {
				defer wg.Done()
				sum := 0
				for k := 0; k < c1; k++ {
					sum += A[i][k] * B[k][j]
				}
				result[i][j] = sum
			}(i, j) // immidiately call the function and pass i and j to the goroutine
		}
	}

	wg.Wait()
	return result
}
