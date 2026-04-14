package main

import (
	"testing"
)

func FuzzDoSomething(f *testing.F) {
	f.Add(3, "test&&")
	f.Fuzz(func(f *testing.T, i int, s string) {
		doSomething(i, s)
	})
}

