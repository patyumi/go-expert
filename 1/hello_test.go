package main

import "testing"

// go test: Para rodar o teste
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, my name is Patrícia"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
