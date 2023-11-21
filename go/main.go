package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

const count = 100000

// const count = 200

func main() {
	ch := make(chan struct{}, count)
	starttime := time.Now()
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

	elapsed := time.Since(starttime)
	fmt.Printf("elapsed time: %s", elapsed)
}

func makeAsync(ch chan<- struct{}, f func()) {
	f()
	ch <- struct{}{}
}

func syncConsoleWrite(id string, duration time.Duration) {
	log.Printf("start with id %s", id)
	time.Sleep(duration)
	log.Printf("stop with id %s", id)
}
