package main

import "testing"

// go test: Para rodar o teste
func TestHello(t *testing.T) {
	// Esses são subtestes
	// Para quando voce quer agrupar testes sobre a mesma coisa
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Patrícia", "english")
		want := "Hello, my name is Patrícia"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, my name is ' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, my name is "

		assertCorrectMessage(t, got, want)
	})

	t.Run("write the phrase in Spanish", func(t *testing.T) {
		got := Hello("Patrícia", "spanish")
		want := "Hola, meu nombre es Patrícia"

		assertCorrectMessage(t, got, want)
	})

	t.Run("write the phrase in French", func(t *testing.T) {
		got := Hello("Patrícia", "french")
		want := "Bonjour, Patrícia"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
