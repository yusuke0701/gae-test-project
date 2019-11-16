package main

import (
	"gae-test-project/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	{
		g := router.Group("/api")
		handler.Comments(g.Group("/commnets"))
	}
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}
