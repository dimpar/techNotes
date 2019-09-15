package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		count("sheep")
		wg.Done()
	}()

	go func() {
		count("stars")
		wg.Done()
	}()

	wg.Wait()
}

func count(s string) {
	for i := 0; i<5; i++ {
		fmt.Println(i, s)
		time.Sleep(time.Millisecond * 500)
	}
}