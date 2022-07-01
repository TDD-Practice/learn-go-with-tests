package main

import "testing"

func TestHello(t *testing.T){
	t.Run("When invoked with a name it says Hello,' name'!",
func(t *testing.T) {
	got := sayHello("Matheus")
	want := engHello + "Matheus!"
	if(got != want) {
		t.Errorf("Wanted %q but got %q", want , got)
	}
})

t.Run("When invoked with an empty string it says 'Hello, person!'",
func(t *testing.T) {
	got := sayHello("")
	want := engHello + "person!"
	
	if (got != want) {
	t.Errorf("Wanted %q but got %q", want , got)
}
})
}