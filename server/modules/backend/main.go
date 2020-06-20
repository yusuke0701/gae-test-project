package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/yusuke0701/gae-test-project/handler"
	"github.com/yusuke0701/goutils/gcp"
)

func main() {
	router := gin.Default()

	// middleware
	router.Use(setEnv())

	// set api routing
	{
		api := router.Group(handler.APIPathPrefix)

		// rest api
		handler.Accounts(api.Group("/accounts"))
		handler.Comments(api.Group("/comments"))
		handler.Tags(api.Group("/tags"))
		handler.Threads(api.Group("/threads"))

		// other api
		handler.SignedURLs(api.Group("/url"))
		handler.Users(api.Group("/users"))
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}

// middlewares

func setEnv() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		gcp.SetTraceID(r)

		c.Next()
	}
}
