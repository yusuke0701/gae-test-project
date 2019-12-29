package main

import (
	"fmt"
	"gae-test-project/connection"
	"gae-test-project/handler"
	_ "gae-test-project/util"
	"log"

	"github.com/gin-gonic/gin"
)

const apiVersion = "v1"

func main() {
	router := gin.Default()
	router.Use(setEnv())
	{
		api := router.Group(fmt.Sprintf("/api/%s", apiVersion))
		handler.Comments(api.Group("/comments"))
		handler.SignedURLs(api.Group("/url"))
	}

	// コネクション
	{
		if err := (connection.DataStore{}).Open(); err != nil {
			log.Fatalf("Failed to connect datastore: %v", err)
		}
		if err := (connection.IAM{}).Open(); err != nil {
			log.Fatalf("Failed to connect iamService: %v", err)
		}
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}
