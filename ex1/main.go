package main

import (
	"os"
)

func main() {
	lang := os.Args[1]
	argument := os.Args[2:]
	println(greeting.detect_language(argument, lang))
}
