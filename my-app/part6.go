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
	go func() { 
		for {
			select {
				case s := <-ch1: new_ch <- s
				case s := <-ch2: new_ch <- s
		        }
	        }
	}()
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

func Generator3(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func Generator4(msg string, quit chan bool) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case ch <- fmt.Sprintf("%s", msg):
			case <-quit:
				fmt.Println("Goroutine done")
				return
			}
		}
	}()
	return ch
}

