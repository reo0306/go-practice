package main

import (
	"fmt"
	"sort"
	"log"
	"os"
)

type MyString string

type User struct {
	Name string
	Age int
}

type Value int

type Speaker interface {
	Speak() error
}

type Dog struct {}

func (d *Dog) Speak() error {
	fmt.Println("BowWow")
	return nil
}

type Cat struct {}

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

	user := User {
		Name: "Bob",
		Age: 18,
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

	defer fmt.Println("6")
	defer fmt.Println("5")
	defer fmt.Println("4")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")

	//doSomething()

}
