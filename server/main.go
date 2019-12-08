package main

import (
	_ "gae-test-project/connection"
	"gae-test-project/handler"
	_ "gae-test-project/util"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	{
		api := router.Group("/api")
		handler.Comments(api.Group("/comments"))
	}
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}
