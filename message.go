package main

import "errors"

type Message struct {
	ID      int
	Content string
}

type repository interface {
	GetMessageByID(id int) (string, error)
	Save(m Message)
	Delete(id int)
}

type service interface {
	GetMessage(id int) string
}

var (
	ErrMessageNotFound = errors.New("requested message not found")
)
