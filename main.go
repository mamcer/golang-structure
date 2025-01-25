package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func getDB() *sql.DB {
	var err error
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cookbook")
	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {
	g := gin.Default()

	r := NewRepository(getDB())
	s := NewService(r)

	Handle(g, s)

	g.Run(":2000")
}
