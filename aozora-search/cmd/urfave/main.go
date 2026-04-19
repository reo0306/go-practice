package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

func listStudentsAsJSON(ctx context.Context) error {
	return nil
}

func listStudents(ctx context.Context) error {
	return nil
}

func cmdList(ctx context.Context, c *cli.Command) error {
	if c.Bool("json") {
		return listStudnetsAsJSON(ctx)
	}
	return listStudents(ctx)
}

func main() {
	app := &cli.Command{}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "config",
			Aliases: []string{"C"},
			Usage: "Load configuration from `FILE`",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name: "list"
			Usage: "list students",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name: "json"
					Usage: "output as JSON",
					Value: false,
				},
			},
			Action: cmdList,
		},
	}
	app.Name = "score"
	app.Usage = "Show student's score"
	app.Run(context.Background(), os.Args)
}
