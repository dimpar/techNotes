package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		count("sheep", 500, c1)
		close(c1)
	}()

	go func() {
		count("fish", 2000, c2)
		close(c2)
	}()

	for {
		select {
		case msg1, ok := <-c1:
			fmt.Printf("msg1 [%s]\n", msg1)
			if !ok {
				c1 = nil // need to set it to nil, otherwise it will be listening
			}
		case msg2, ok := <-c2:
			fmt.Printf("msg2 [%s]\n", msg2)
			if !ok {
				c2 = nil
			}
		}

		// if not adding this break, you will enter an infinite loop
		if c1 == nil && c2 == nil {
			break
		}
	}
}

func count(s string, interval time.Duration, c chan<- string) {
	for i := 0; i < 10; i++ {
		c <- s
		time.Sleep(interval * time.Millisecond)
	}
}
