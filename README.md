
# Matrix Multiplication in Go — With and Without Goroutines

This project demonstrates matrix multiplication implemented in Go using two approaches:

- Standard (sequential) multiplication
- Concurrent multiplication using goroutines

It is designed to help understand the concept of threads and concurrency by comparing a simple implementation with a concurrent one using Go’s powerful goroutine model.


# Note:


- Each matrix element computation runs in its own goroutine.

- How do we know when all the goroutines have finished?

        A sync.WaitGroup ensures all computations complete before proceeding.
        - wg.Add(1) adds to the counter
        - wg.Done() decrements it when a goroutine finishes
        - wg.Wait() blocks until all tasks are done

        So, if we have a 500*500 matrix, the sequence will be:
                1. wg.Add(1) runs 250000 times (the internal WaitGroup counter goes to 250,000).
                2. the go func spawns 250,000 goroutines.
                3. inside each goroutine, defer wg.Done() is called at the end (gradually decrementing the counter by 1)
                4. wg.Wait() doesnt let the main exit until the counter becomes 0

- Goroutines + WaitGroup follow the fork-join principle  
  - **Fork:** Each matrix cell calculation is forked into a separate goroutine.  
  - **Join:** All goroutines are joined using a sync.WaitGroup.



- You can observe concurrency by adding a debug print inside the goroutine. If you do so, you will notice the output seems random (i.e not from top-left to bottom-right order)

- To finally measure, how much faster the goroutine-based approach is compared to the standard sequential method, i timed both the implementations on the same 500*500 matrix.

> **Results:**
> - Sequential: 376.808046ms  
> - Parallel:   110.450523ms

- Caution: For a large matrix (n*n), this program will launch n^2 goroutines. Each goroutine consumes memory and scheduling time. Hence, this is not scalable for much larger matrices. A better approach could be limiting the number of goroutines to 1 per row instead of 1 per cell. This balances performance and memory usage better on large-scale tasks.


  

  ## Run

- **go run main.go**  
  Runs the interactive matrix multiplication program. You’ll be prompted to enter matrix dimensions and values.

- **go run compare/compare.go**
  Benchmarks sequential vs goroutine-based multiplication using the same randomly generated 500×500 matrix.


## Useful Resources:


- [Golang Concurrency Guide – GoLinuxCloud](https://www.golinuxcloud.com/golang-concurrency/)
- [ Go Routine Internals without breaking your brain  ](https://www.youtube.com/watch?v=qyM8Pi1KiiM)
- [ Master Go Programming With These Concurrency Patterns (in 40 minutes) ](https://www.youtube.com/watch?v=O4aTkRQAK-o)
  



