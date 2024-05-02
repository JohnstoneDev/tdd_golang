package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	// Helper function that helps in printing an Error message
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Darth Vader", "Spanish")
		want := "Hola, Darth Vader"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in English", func(t *testing.T){
		got := Hello("Darth Vader", "English")
		want := "Hello, Darth Vader"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T){
		got := Hello("Darth Vader", "French")
		want := "Bonjour, Darth Vader"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}