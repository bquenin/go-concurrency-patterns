package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Event struct {
	value int
}

func generateEvents() <-chan Event {
	events := make(chan Event)

	go func() {
		for i := 0; ; i++ {
			events <- Event{i}
			fmt.Println("event", i, "published")
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()

	return events
}

func bufferEvents(input <-chan Event, duration time.Duration) <-chan []Event {
	buffers := make(chan []Event)

	go func() {
		var buffer []Event
		ticker := time.Tick(duration)
		for {
			select {
			case <-ticker:
				buffers <- buffer
				buffer = nil
			case event := <-input:
				buffer = append(buffer, event)
			}
		}
	}()

	return buffers
}

func main() {
	rand.Seed(time.Now().UnixNano())

	events := generateEvents()

	eventBuffers := bufferEvents(events, 5*time.Second)
	for eventBuffer := range eventBuffers {
		fmt.Println("events", eventBuffer)
	}
}
