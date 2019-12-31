package util

import (
	"cloud.google.com/go/datastore"
	"google.golang.org/api/iam/v1"
)

// DatastoreClient は、データストアへの接続を担保する
var DatastoreClient *datastore.Client

// IAMService は、IAMサービスへの接続を担保する
var IAMService *iam.Service
