package main

type Message struct {
	ID      int
	Content string
}

type repository interface {
	GetMessageByID(id int) string
	Save(m Message)
	Delete(id int)
}

type service interface {
	GetMessage(id int) string
}
