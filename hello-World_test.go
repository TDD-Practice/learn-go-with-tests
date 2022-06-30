package main

import "testing"

func TestHelloWorld(t *testing.T){
	got := HelloWorld()
	want := "Hola Matheus!"
	
	if got != want {
		t.Errorf("Got %q but I wanted %q", got, want)
	}
}