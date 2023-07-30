package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func numberDifferentIntegers(word string) int {
	integerBucket := map[int]struct{}{}
	tmp_number := ""
	for _, char := range word {
		if unicode.IsDigit(char) {
			tmp_number += string(char)
		} else {
			if tmp_number != "" {
				number, _ := strconv.Atoi(tmp_number)
				integerBucket[number] = struct{}{}
				tmp_number = ""
			}
		}
	}
	return len(integerBucket)
}

func PublicCountInteger(word string) int {
	return numberDifferentIntegers(word)
}

func TryingCountInteger() {
	word := "a123bc34d8ef34"
	count := numberDifferentIntegers(word)
	fmt.Println(count)
	word2 := "A1b01c001"
	count2 := numberDifferentIntegers(word2)
	fmt.Println(count2)
}
