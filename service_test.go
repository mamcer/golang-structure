package main

import "testing"

type MockRepository struct {
}

func (*MockRepository) GetMessageByID(id int) string {
	return "pong"
}

func (*MockRepository) Save(m Message) {}

func (*MockRepository) Delete(id int) {}

func TestGetMessage(t *testing.T) {
	t.Run("get message from service", func(t *testing.T) {
		s := NewService(&MockRepository{})
		got := s.GetMessage(1)
		want := "pong"
		assertCorrectMessage(t, got, want)
	})
}
