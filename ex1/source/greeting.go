package ex1

import (
	"fmt"
	"log"
	"os"

	"github.com/tuanhuu162/go23_course/ex1/constants"
	"github.com/urfave/cli/v2"
)

func english_formating(argument []string) string {
	if len(argument) == 2 {
		first_name := argument[0]
		last_name := argument[1]
		return fmt.Sprintf("Hello dear %q %q", first_name, last_name)
	} else if len(argument) == 3 {
		first_name := argument[0]
		middle_name := argument[1]
		last_name := argument[2]
		return fmt.Sprintf("Hello dear %q %q %q", first_name, middle_name, last_name)
	} else {
		return fmt.Sprintf("Hello dear %q", argument[0])
	}
}

func vietnamese_foramting(argument []string) string {
	if len(argument) == 2 {
		last_name := argument[0]
		first_name := argument[1]
		return fmt.Sprintf("Xin chao %q %q", last_name, first_name)
	} else if len(argument) == 3 {
		first_name := argument[2]
		middle_name := argument[1]
		last_name := argument[1]
		return fmt.Sprintf("Xin chao %q %q %q", last_name, middle_name, first_name)
	} else {
		return fmt.Sprintf("Xin chao %q", argument[0])
	}
}

func detect_language(argument []string, language constants.LANGUAGE) string {
	switch language {
	case constants.English:
		return english_formating(argument)
	case constants.Vietnamese:
		return vietnamese_foramting(argument)
	default:
		panic(fmt.Errorf("unsupported language: %s", language))
	}
}

func hello() {
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

				if language == "vietnamese" {
					fmt.Printf("Xin chao %q", name)
				} else {
					fmt.Printf("Hello dear %q", name)
				}
				return nil
			}
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
