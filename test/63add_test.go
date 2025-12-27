package main

import (
	"testing"
)

func TestAdd(t *testing.T) {

	result := add(5, 18)
	expected := 23
	if result != expected {
		t.Errorf("Add(5,18) failed: expected is %v but got %v\n ", expected, result)
	}
}
