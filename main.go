package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func getDB(dbdriver string, dbsource string) *sql.DB {
	var err error
	db, err := sql.Open(dbdriver, dbsource)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {
	port := os.Getenv("FLAT_PORT")
	dbdriver := os.Getenv("FLAT_DB_DRIVER_NAME")
	dbsource := os.Getenv("FLAT_DB_DATA_SOURCE_NAME")

	g := gin.Default()

	r := NewRepository(getDB(dbdriver, dbsource))
	s := NewService(r)

	Handle(g, s)

	g.Run(":" + port)
}
