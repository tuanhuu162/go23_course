package main

import (
	"testing"

	"github.com/go23_course/ex3_and_ex4"
)

func TestCountIntegerNormalCase(t *testing.T) {
	testInput := "a123bc34d8ef34"
	expect_result := 3
	result := ex3_and_ex4.PublicCountInteger(testInput)
	if result != expect_result {
		t.Errorf("Expect there are %d rectangles, but got %d", expect_result, result)
	}
}
