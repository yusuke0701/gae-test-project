package connection

import (
	"context"

	"cloud.google.com/go/datastore"
)

// DatastoreClient は、データストアへの接続を担保する
var DatastoreClient *datastore.Client

type DataStore struct{}

func (d DataStore) Open(projectID string) (err error) {
	ctx := context.Background()
	DatastoreClient, err = datastore.NewClient(ctx, projectID)
	return
}
