package main

import (
	"errors"
	"testing"
)

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGetMessageByID(t *testing.T) {
	t.Run("testing pong return, message id 1", func(t *testing.T) {
		r := NewRepository(nil)
		got, e := r.GetMessageByID(1)
		if e == nil {
			want := "pong"
			assertCorrectMessage(t, got, want)
		} else {
			t.Errorf("expect error nil, got:'%v'", e)
		}
	})

	t.Run("testing error, message id 2", func(t *testing.T) {
		r := NewRepository(nil)
		got, e := r.GetMessageByID(2)

		if !errors.Is(e, ErrMessageNotFound) {
			t.Errorf("expect error '%v', got:'%v'", ErrMessageNotFound, e)
		} else {
			want := ""
			assertCorrectMessage(t, got, want)
		}

	})
}
