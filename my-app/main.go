package main

import (
	"log"
	"log/slog"
	"os"
	"time"

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
}
