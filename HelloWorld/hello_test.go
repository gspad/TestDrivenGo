package main

import "testing"

func TestHelloWorld(t *testing.T) {
	//given a name
	t.Run("when a name is provided, say hello to them", func(t *testing.T) {
		got := Hello("Gianluca", "English")
		want := "Hello, Gianluca"

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})

	t.Run("when empty string is passed, default to returning Hello, world", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})

	t.Run("when given a name and language is English, return a salute in English", func(t *testing.T) {
		got := Hello("Michelle", "English")
		want := "Hello, Michelle"

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})

	t.Run("when given a name and language is Spanish, return a salute in Spanish", func(t *testing.T) {
		got := Hello("Michelle", "Spanish")
		want := "Hola, Michelle"

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})

	t.Run("when given a name and the language is French, return a salute in French", func(t *testing.T) {
		got := Hello("Michelle", "French")
		want := "Bonjour, Michelle"

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})
}
