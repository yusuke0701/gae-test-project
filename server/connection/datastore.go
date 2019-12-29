package connection

import (
	"context"
	"gae-test-project/util"

	"cloud.google.com/go/datastore"
)

// DatastoreClient は、データストアへの接続を担保する
var DatastoreClient *datastore.Client

type DataStore struct{}

func (d DataStore) Open() (err error) {
	ctx := context.Background()
	DatastoreClient, err = datastore.NewClient(ctx, util.ProjectID)
	return
}
