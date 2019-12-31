package main

import (
	"fmt"
	"gae-test-project/connection"
	"gae-test-project/handler"
	"gae-test-project/util"
	"log"

	"github.com/gin-gonic/gin"
)

const apiVersion = "v1"

func main() {
	// connection
	{
		if err := (connection.DataStore{}).Open(util.ProjectID); err != nil {
			log.Fatalf("Failed to connect datastore: %v", err)
		}
		if err := (connection.IAM{}).Open(); err != nil {
			log.Fatalf("Failed to connect iamService: %v", err)
		}
	}
	// routing
	{
		router := gin.Default()

		// middleware
		router.Use(setEnv())

		{
			api := router.Group(fmt.Sprintf("/api/%s", apiVersion))
			handler.Comments(api.Group("/comments"))
			handler.SignedURLs(api.Group("/url"))
		}

		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
	}
}
