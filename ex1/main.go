package main

import (
	"fmt"
	"os"

	"github.com/tuanhuu162/go23_course/ex1/source"
)

func main() {
	lang := os.Args[1]
	argument := os.Args[2:]
	greeting := source.Greeting{}
	var result = greeting.DetectLanguage(argument, lang)
	fmt.Println(result)
}
