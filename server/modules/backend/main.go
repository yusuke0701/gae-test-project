package main

import (
	"context"
	"gae-test-project/firebase"
	"gae-test-project/handler"
	"gae-test-project/util"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iam/v1"
)

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

			if err := firebase.Setup(ctx, "AIzaSyCEdlcKinO_em8f_ymWrE3_qAkaMLftNms"); err != nil {
				log.Fatalf("Failed to connect firebase: %v", err)
			}
		}
	}
	// routing
	{
		router := gin.Default()

		// middleware
		router.Use(setEnv())

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
}

// middlewares

func setEnv() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		util.SetTraceID(r)

		c.Next()
	}
}
