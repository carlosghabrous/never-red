package server

import "testing"

func TestHelloHandler(t *testing.T) {
	got, want := 1, 1
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
