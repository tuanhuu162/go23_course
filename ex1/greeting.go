package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func greeting_without_lib() {
	var language string
	if len(os.Args) > 1 {
		language = os.Args[1]
	}
	name := "unknown"
	if len(os.Args) > 2 {
		name = os.Args[2]
	}
	if language == "japanese" {
		fmt.Printf("こんにちは %q です。", name)
	} else if language == "vietnamese" {
		fmt.Printf("Xin chao %q", name)
	} else {
		fmt.Printf("Hello dear %q", name)
	}
}

func greeting() {
	var language string

	app := &cli.App{
		Name:  "greeting",
		Usage: "Personalized greeting command line tool for newbie with go",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang",
				Value:       "english",
				Usage:       "language for the greeting",
				Destination: &language,
			},
		},
		Action: func(context *cli.Context) error {
			name := "unknown"
			if context.NArg() > 0 {
				name = context.Args().Get(0)
			}
			if language == "japanese" {
				fmt.Printf("こんにちは %q です。", name)
			} else if language == "vietnamese" {
				fmt.Printf("Xin chao %q", name)
			} else {
				fmt.Printf("Hello dear %q", name)
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
