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

func FanIn(ch1, ch2 <-chan string) <-chan string {
	new_ch := make(chan string)
	go func() { for { new_ch <- <-ch1 } } ()
	go func() { for { new_ch <- <-ch2 } } ()
	return new_ch
}

func Generator2(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}

