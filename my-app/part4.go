package main

import (
	"fmt"
	//"os"
)

type F struct {
	Name string
	Age int
}

func (f *F) String() string {
	return fmt.Sprintf("NAME=%q, AGE=%d", f.Name, f.Age)
}

func Part4() {
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
