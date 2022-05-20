package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("tkm")
	want := "Hello, tkm"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
