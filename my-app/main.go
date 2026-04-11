package main

import (
	"log"
	"log/slog"
	"os"
	"time"
	"fmt"
	"math/rand"

	"my-app/server"
)

var name = "John"

func init() {
	println("Hi! " + name)
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))

}

func main() {
	Part3()
	Part4()


	f, err := os.Create("server.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	svr := server.New("localhost", 8080,
		server.WithTimeout(time.Minute),
		server.WithLogger(logger),
	)
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}

	Part5()

	ch := Generator("Hello")
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}

	ch2 := FanIn(Generator2("Hello"), Generator2("Byb"))
	for i := 0; i < 10; i++ {
		fmt.Println(<- ch2)
	}

	ch3 := Generator3("Hi")
	for i := 0; i < 10; i++ {
		select {
		case s := <-ch3:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("Waited too long!")
			return
		}
	}

	quit := make(chan bool)
	ch4 := Generator4("Hi!", quit)
	for i := rand.Intn(50); i >= 0; i-- {
		fmt.Println(<-ch4, i)
	}
	quit <- true
}
