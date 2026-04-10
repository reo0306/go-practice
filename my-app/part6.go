package main

import (
	"fmt"
	"time"
)

func Generator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg ,i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}
