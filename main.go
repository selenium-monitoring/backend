package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	apiGroup := r.Group("/api")

	apiGroup.GET("/ping", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	r.Run()
}
