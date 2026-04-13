package main

import "time"

func StringDistance(lhs, rhs string) int {
	return Distance([]rune(lhs), []rune(rhs))
}

func Distance(a, b []rune) int {
	dist := 0
	if len(a) != len(b) {
		return -1
	}

	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist
}


func Add(a, b int) int {
	result := a + b
	time.Sleep(3 * time.Second)
	return result
}
