package main

import (
	"fmt"
	//"os"
	"time"
	"sync"
	"context"
	"iter"
)

type F struct {
	Name string
	Age int
}

func (f *F) String() string {
	return fmt.Sprintf("NAME=%q, AGE=%d", f.Name, f.Age)
}

func ff(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		fmt.Println("goroutine: 処理")
		time.Sleep(1 * time.Second)
	}
}

func iter1[T any](a []T) iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, v := range a {
			if !yield(fmt.Sprint(v)) {
				break
			}
		}
	}
}

func Part4() {
	for v := range iter1([]int{4,5,6}) {
		println(v)
	}

	iter1 := func(yield func(string) bool) {
		a := []int{1, 2, 3}
		for _, v := range a {
			if !yield(fmt.Sprint(v)) {
				break
			}
		}
	}

	for v := range iter1 {
		println(v)
	}

	f := &F{
		Name: "John",
		Age: 20,
	}
	
	fmt.Printf("%v\n", f)
	fmt.Printf("%#v\n", f)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", *f)

	println("Hello! " + name)

	x := 1
	s := fmt.Sprint(x)
	fmt.Println(s)


	now := time.Now()
	fmt.Println(now.Format(time.RFC3339))


	var wg sync.WaitGroup
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go ff(ctx, &wg)

	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()

	//f, err := os.Create("output.dat")
	//if err != nil {
		//panic(err)
	//}
	//defer f.Close()

	//fmt.Fprintf(f, "COUNT %05d\n", 5)


	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("%[1]T: %[1]s\n", e)
		}
	}()

	panic("my error")
	//var a [2]int
	//n := 2
	//println(a[n])
}
