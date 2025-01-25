package main

import "testing"

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGetMessageByID(t *testing.T) {
	t.Run("testing pong return", func(t *testing.T) {
		r := NewRepository(nil)
		got := r.GetMessageByID(1)
		want := "pong"
		assertCorrectMessage(t, got, want)
	})
}
