package main

import "testing"

func TestHelloYou (t *testing.T) {
	got := helloYou("Matheus")
	want := "Hello Matheus!"

	if(got != want){
		t.Errorf("I wanted %q but I got %q", want, got)
	}
}