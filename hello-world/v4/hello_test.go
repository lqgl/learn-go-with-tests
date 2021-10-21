package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chirs")
	want := "Hello, Chirs"

	if got != want {
		t.Errorf("got '%q', but want '%q'", got, want)
	}
}
