package main

import "database/sql"

type Repository struct {
	Conn *sql.DB
}

func NewRepository(conn *sql.DB) *Repository {
	return &Repository{conn}
}

func (*Repository) GetMessageByID(id int) (string, error) {
	if id == 1 {
		return "pong", nil
	}

	return "", ErrMessageNotFound
}

func (*Repository) Save(m Message) {
}

func (*Repository) Delete(id int) {
}
