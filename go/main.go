package main

import (
	"log"
	"strconv"
	"time"
)

const count = 100000

func main() {
	ch := make(chan struct{}, count+1)
	for i := 0; i < count; i++ {
		id := strconv.Itoa(i)
		go makeAsync(ch, func() {
			syncConsoleWrite(id, 1*time.Second)
		})
	}

	for i := 0; i < count; i++ {
		<-ch

		if i%1000 == 0 {
			log.Printf("iteration %s", strconv.Itoa(i))
		}
	}
}

func makeAsync(ch chan<- struct{}, f func()) {
	f()
	ch <- struct{}{}
}

func syncConsoleWrite(id string, duration time.Duration) {
	// log.Printf("start with id %s", id)
	time.Sleep(duration)
	// log.Printf("stop with id %s", id)
}
