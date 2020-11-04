package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		count("sheep", 500)
		wg.Done()
	}()

	go func() {
		count("fish", 2000)
		wg.Done()
	}() 
	
	wg.Wait()
}

func count(s string, interval time.Duration) {
	for i := 0; i < 10; i++ {
		fmt.Println(s)
		time.Sleep(interval * time.Millisecond)
	}
}