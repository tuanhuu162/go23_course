package main

import (
	"testing"
)

func TestCountIntegerNormalCase(t *testing.T) {
	testInput := "a123bc34d8ef34"
	expect_result := 3
	result := PublicCountInteger(testInput)
	if result != expect_result {
		t.Errorf("Expect there are %d rectangles, but got %d", expect_result, result)
	}
}

func TestCountIntegerEmptyString(t *testing.T) {
	testInput := ""
	expect_result := 0
	result := PublicCountInteger(testInput)
	if result != expect_result {
		t.Errorf("Expect there are %d rectangles, but got %d", expect_result, result)
	}
}

func TestCountIntegerNotFoundInteger(t *testing.T) {
	testInput := "abcdef"
	expect_result := 0
	result := PublicCountInteger(testInput)
	if result != expect_result {
		t.Errorf("Expect there are %d rectangles, but got %d", expect_result, result)
	}
}
