package store

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/yusuke0701/goutils/gcp"
)

// datastoreClient は、データストアへの接続を担保する
var datastoreClient *datastore.Client

func init() {
	ctx := context.Background()

	dc, err := datastore.NewClient(ctx, gcp.ProjectID)
	if err != nil {
		log.Fatalf("Failed to connect datastore: %v", err)
	}
	datastoreClient = dc
}
