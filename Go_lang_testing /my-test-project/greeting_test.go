package main

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting("Chris")
	want := "Hello, chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
