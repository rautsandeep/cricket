package main

import(
"testing"
"log"
)

func TestIspalindrome(t *testing.T){
	got:= ispalindrome([]byte("tyuiuyt"))
	want:= true

	if got != want {
		log.Fatalf("error: test failed  %v is not equat to %v\n",got,want)
	}

}
