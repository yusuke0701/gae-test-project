package connection

import (
	"context"
	"gae-test-project/util"
	"log"

	"cloud.google.com/go/datastore"
)

// DatastoreClient は、データストアへの接続を担保する
var DatastoreClient *datastore.Client

func init() {
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, util.ProjectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	DatastoreClient = client
}
