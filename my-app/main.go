package main

import "fmt"
import "sort"

type MyString string

type User struct {
	Name string
	Age int
}

type Value int

func showName(user *User) {
	fmt.Println(user.Name)
}

func (v Value) Add(n Value) Value {
	return v + n
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
	v = v.Add(2)
	fmt.Println(v)

	s := 1
	p := &s
	*p = 2
	fmt.Println(s)


}
