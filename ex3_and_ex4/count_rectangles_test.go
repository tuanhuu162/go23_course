package main

import (
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
	result := PublicCountRectangles(testInput)
	if result != expect_result {
		t.Errorf("Expect there are %d rectangles, but got %d", expect_result, result)
	}
}

func TestCountRectabglesEmptyArray(t *testing.T) {
	testInput := [][]int{}
	expect_result := 0
	result := PublicCountRectangles(testInput)
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
	result := PublicCountRectangles(testInput)
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
	result := PublicCountRectangles(testInput)
	if result != expect_result {
		t.Errorf("Expect there are %d rectangles, but got %d", expect_result, result)
	}
}
