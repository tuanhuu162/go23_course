package source

import (
	"fmt"
	"log"
	"os"

	"github.com/tuanhuu162/go23_course/ex1/constants"
	"github.com/urfave/cli/v2"
)

type Greeting struct {
}

func english_formating(argument []string) string {
	if len(argument) == 2 {
		first_name := argument[0]
		last_name := argument[1]
		return fmt.Sprintf("Hello dear %s %s", first_name, last_name)
	} else if len(argument) == 3 {
		first_name := argument[0]
		middle_name := argument[1]
		last_name := argument[2]
		return fmt.Sprintf("Hello dear %s %s %s", first_name, middle_name, last_name)
	} else {
		return fmt.Sprintf("Hello dear %s", argument[0])
	}
}

func vietnamese_foramting(argument []string) string {
	if len(argument) == 2 {
		last_name := argument[1]
		first_name := argument[0]
		return fmt.Sprintf("Xin chao %s %s", last_name, first_name)
	} else if len(argument) == 3 {
		first_name := argument[0]
		middle_name := argument[1]
		last_name := argument[2]
		return fmt.Sprintf("Xin chao %s %s %s", last_name, middle_name, first_name)
	} else {
		return fmt.Sprintf("Xin chao %s", argument[0])
	}
}

func (f Greeting) DetectLanguage(argument []string, language string) string {
	switch constants.LANGUAGE(language) {
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
					fmt.Printf("Xin chao %s", name)
				} else {
					fmt.Printf("Hello dear %s", name)
				}
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
