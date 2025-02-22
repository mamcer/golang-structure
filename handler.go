package main

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service
}

func (ph *Handler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": ph.Service.GetMessage(1),
	})
}

func Handle(g *gin.Engine, ps service) {
	ph := Handler{ps}
	g.GET("/ping", ph.Ping)
}
