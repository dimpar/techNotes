package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go count("sheep", c)

	for {
		msg, open := <- c //can receive a second arg, which is open
		if !open {
			break
		}
		fmt.Println(msg)
	}

	// this is the same as above, just a syntatic sugar
	for msg1 := range c {
		fmt.Println(msg1)
	}

}

func count(s string, c chan string) {
	for i := 0; i<5; i++ {
		c <- s // Sending to a channel is a blocking operation
		fmt.Println(i)
		time.Sleep(time.Millisecond * 500)
	}

	close(c) // Note: if you are a receiver, you should never close a channel. Only a receiver should close a channel.
}