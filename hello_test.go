package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("say hello to the named person", func(t *testing.T) {

		got := Hello("Meera")
		want := "Hello, Meera"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, World' when an empty string is passed", func(t *testing.T) {

		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
