package main

import (
	"fmt"
	"os"

	"mmake/commands"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "mmake",
		Usage: "a quick make alternative intended for personal purposes",
		Action: func(ctx *cli.Context) error {
			_, err := commands.Compile(ctx)

			return err
		},
		Commands: []*cli.Command{
			{
				Name:    "make",
				Aliases: []string{"compile", "m", "c"},
				Action: func(ctx *cli.Context) error {
					_, err := commands.Compile(ctx)

					return err
				},
			},
			{
				Name:    "clean",
				Aliases: []string{"clean", "fix", "tidy"},
				Action: func(ctx *cli.Context) error {
					return commands.Clean(ctx)
				},
			},
			{
				Name:    "run",
				Aliases: []string{"r"},
				Action: func(ctx *cli.Context) error {
					return commands.Run(ctx)
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "args",
						Aliases: []string{"a", "arg", "arguments"},
						Usage:   "arguments passed to the run program",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("error: %s", err.Error())
	}
}
