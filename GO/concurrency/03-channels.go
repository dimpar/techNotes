package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go count("sheep", 500, c)
	
	for msg := range c {
		fmt.Printf("msg [%s]\n", msg)
	}
}

func count(s string, interval time.Duration, c chan string) {
	for i := 0; i < 10; i++ {
		c <- s
		time.Sleep(interval * time.Millisecond)
	}

	close(c)
}
