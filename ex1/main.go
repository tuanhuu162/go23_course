package main

import (
	"os"
)

func main() {
	lang := os.Args[1]
	argument := os.Args[2:]
	println(source.detect_language(argument, lang))
}
