package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	//"time"
	"sync"

	"github.com/mattn/go-runewidth"
)

type MyString string

type User struct {
	Name string
	Age  int
}

type Value int

type Speaker interface {
	Speak() error
}

type Dog struct{}

func (d *Dog) Speak() error {
	fmt.Println("BowWow")
	return nil
}

type Cat struct{}

func (c *Cat) Speak() error {
	fmt.Println("Meow")
	return nil
}

func DoSpeak(s Speaker) error {
	return s.Speak()
}

func showName(user *User) {
	fmt.Println(user.Name)
}

func (v *Value) Add(n Value) {
	*v += n
}

func doSomething() {
	var n = 1
	defer func() {
		fmt.Println(n)
	}()

	n = 2
}

func sendMessage(msg string) {
	println(msg)
}

func server(ch chan string) {
	defer close(ch)
	ch <- "one"
	ch <- "two"
	ch <- "three"
}

func main() {
	m := make(map[string]int)
	m["John"] = 23
	m["Bob"] = 21
	m["Mark"] = 22
	fmt.Println(m)

	for k, v := range m {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("key: %v, value: %v", k, m[k])
	}

	var a MyString
	a = "foo"
	fmt.Println(a)

	user := User{
		Name: "Bob",
		Age:  18,
	}

	showName(&user)

	v := Value(1)
	v.Add(2)
	fmt.Println(v)

	s := 1
	p := &s
	*p = 2
	fmt.Println(s)

	var j interface{}
	j = 1
	n := j.(int)
	fmt.Println(n)

	j = "Hello world"
	b := j.(string)
	fmt.Println(b)

	dog := Dog{}
	DoSpeak(&dog)

	cat := Cat{}
	DoSpeak(&cat)

	f, err := os.Open("./data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var k [512]byte
	nn, err := f.Read(k[:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(k[:nn]))

	//defer fmt.Println("6")
	//defer fmt.Println("5")
	//defer fmt.Println("4")
	//fmt.Println("1")
	//fmt.Println("2")
	//fmt.Println("3")

	//doSomething()

	//message := "hi"
	//go func() {
	//sendMessage(message)
	//}()

	//message = "ho"

	//time.Sleep(time.Second)
	//println(message)
	//time.Sleep(time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		vv := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(vv)
		}()
	}
	wg.Wait()

	aa := 0

	var mu sync.Mutex

	var wg2 sync.WaitGroup
	wg2.Add(2)

	go func() {
		defer wg2.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			aa++
			mu.Unlock()
		}
	}()

	go func() {
		defer wg2.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			aa++
			mu.Unlock()
		}
	}()

	wg2.Wait()
	fmt.Println(aa)

	var bb string
	ch := make(chan string)
	go server(ch)

	bb = <-ch
	fmt.Println(bb)
	bb = <-ch
	fmt.Println(bb)
	bb = <-ch
	fmt.Println(bb)

	fmt.Println(runewidth.StringWidth("こんにちは"))

}
