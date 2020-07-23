package store

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"

	"github.com/yusuke0701/goutils/gcp"
)

var datastoreClient *datastore.Client

func init() {
	ctx := context.Background()

	projectID, _ := gcp.GetGAEVar()

	dc, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to connect datastore: %v", err)
	}

	datastoreClient = dc
}
