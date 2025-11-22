package main

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	text string
}

func main() {
	c := make(chan Message)

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()

		m := <-c

		fmt.Println(m)
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c <- Message{text: "Hello World"}
	}()

	wg.Wait()
}
