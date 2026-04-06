package main

import "fmt"

type Fruit int

const (
	Apple Fruit = iota
	Orange
	Banana
)

func (i Fruit) String() string {
	switch i {
	case Apple:
		return "Apple"
	case Orange:
		return "Orage"
	case Banana:
		return "Banana"
	}

	return fmt.Sprintf("Fruit(%d)", i)
}

