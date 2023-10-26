package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/api")
	apiGroup.GET("/ping", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
