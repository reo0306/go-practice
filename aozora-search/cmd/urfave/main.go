package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "config",
			Aliases: []string{"C"},
			Usage: "Load configuration from `FILE`",
		},
	}
	app.Name = "score"
	app.Usage = "Show student's score"
	app.Run(context.Background(), os.Args)
}
