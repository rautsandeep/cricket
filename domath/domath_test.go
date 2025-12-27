package main

import (
	"testing"
)

func TestDomath(t *testing.T) {
	result := domath(5, 9, add)
	expected := 14

	if result != expected {
		t.Errorf("expected %v and got %v", expected, result)
	}

}
