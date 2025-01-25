package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service
}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}

func (ph *Handler) Ping(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	c.JSON(200, gin.H{
		"message": ph.Service.GetMessage(1),
	})
}

func Handle(g *gin.Engine, ps service) {
	ph := Handler{ps}
	g.GET("/ping", ph.Ping)
	g.OPTIONS("/ping", preflight)
}
