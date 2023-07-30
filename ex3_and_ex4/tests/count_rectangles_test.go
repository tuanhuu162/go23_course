package main

import (
	"github.com/tuanhuu162/go23_course/ex3_and_ex4/count_rectangles"
	"testing"
)

func TestCountRectabglesNormalCase(t *testing.T) {
	testInput := [][]int{
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}
	expect_result := 5
	result := count_rectangles.PublicCountRectangles(testInput)
	if result != expect_result {
		t.Errorf("Expect there are %d rectangles, but got %d", expect_result, result)
	}

	// Another test case
	testInput2 := [][]int{
		{1, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}
	expect_result2 := 7
	result2 := count_rectangles.PublicCountRectangles(testInput2)
	if result != expect_result {
		t.Errorf("Expect there are %d rectangles, but got %d", expect_result2, result2)
	}
}

func TestCountRectabglesEmptyArray(t *testing.T) {
	testInput := [][]int{}
	expect_result := 0
	result := count_rectangles.PublicCountRectangles(testInput)
	if result != expect_result {
		t.Errorf("Expect there are %d rectangles, but got %d", expect_result, result)
	}
}

func TestCountRectabglesNWithoutRectangle(t *testing.T) {
	testInput := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
	expect_result := 0
	result := count_rectangles.PublicCountRectangles(testInput)
	if result != expect_result {
		t.Errorf("Expect there are %d rectangles, but got %d", expect_result, result)
	}
}

func TestCountRectabglesAdjacentCase(t *testing.T) {
	testInput := [][]int{
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}
	expect_result := 4
	result := count_rectangles.PublicCountRectangles(testInput)
	if result != expect_result {
		t.Errorf("Expect there are %d rectangles, but got %d", expect_result, result)
	}
}
