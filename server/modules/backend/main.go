package main

import (
	"context"
	"fmt"
	"gae-test-project/handler"
	"gae-test-project/util"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iam/v1"
)

const apiVersion = "v1"

func main() {
	// connection
	{
		ctx := context.Background()

		dc, err := datastore.NewClient(ctx, util.ProjectID)
		if err != nil {
			log.Fatalf("Failed to connect datastore: %v", err)
		}
		util.DatastoreClient = dc

		if !util.IsLocal {
			is, err := iam.NewService(ctx)
			if err != nil {
				log.Fatalf("Failed to connect iamService: %v", err)
			}
			util.IAMService = is
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

// middlewares

func setEnv() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		util.SetTraceID(r)

		c.Next()
	}
}
