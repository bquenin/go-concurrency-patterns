package main

import (
	"fmt"
	"sync"
	"time"
)

const numberOfTasks = 1000
const maximumConcurrentTasks = 10

func task(id int) {
	fmt.Println("task", id, "starting")
	time.Sleep(time.Second)
	fmt.Println("task", id, "completed")
}

func main() {
	semaphore := make(chan struct{}, maximumConcurrentTasks)

	var wg sync.WaitGroup
	for i := 0; i < numberOfTasks; i++ {
		wg.Add(1)
		go func(taskId int) {
			semaphore <- struct{}{}
			defer wg.Done()
			task(taskId)
			<-semaphore
		}(i)
	}
	wg.Wait()
}
