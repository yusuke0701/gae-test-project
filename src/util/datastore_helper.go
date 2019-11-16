package util

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/datastore"
)

// DatastoreClient は、データストアへの接続を担保する
var DatastoreClient *datastore.Client

func init() {
	ctx := context.Background()
	projectID := os.Getenv("PROJECT_ID")

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	DatastoreClient = client
}
