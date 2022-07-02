package main

import (
	"fmt"
	"math/rand"
	"time"
)

const numberOfWorkers = 10

func jobsGenerator() <-chan int {
	jobs := make(chan int)
	go func() {
		for i := 0; ; i++ {
			jobs <- rand.Int()
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return jobs
}

type result struct {
	val      int
	workerId int
}

func startWorkerPool(jobs <-chan int, poolSize int) <-chan result {
	results := make(chan result)
	for i := 0; i < poolSize; i++ {
		go func(workerId int) {
			for job := range jobs {
				fmt.Println("job", job, "assigned to worker", workerId)
				results <- result{job / 2, workerId}
			}
		}(i)
	}
	return results
}

func main() {
	jobs := jobsGenerator()
	results := startWorkerPool(jobs, numberOfWorkers)
	for result := range results {
		fmt.Println("result value", result.val, "computed by worker", result.workerId)
	}
}
