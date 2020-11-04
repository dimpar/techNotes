package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep", 500)
	count("fish", 2000)
}

func count(s string, interval time.Duration) {
	for {
		fmt.Println(s)
		time.Sleep(interval * time.Millisecond)
	}
}