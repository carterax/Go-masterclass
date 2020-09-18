package main

import (
	"testing"
)

func TestCheckEvenOrOdd(t *testing.T) {
	num := 3
	res := checkEvenOrOdd([]int{num})
	if res != "3 is odd" {
		t.Errorf("Expected %v is odd instead got %v is even", num, num)
	}
}
