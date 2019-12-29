package connection

import (
	"context"

	"google.golang.org/api/iam/v1"
)

// IAMService は、IAMサービスへの接続を担保する
var IAMService *iam.Service

type IAM struct{}

func (i IAM) Open() (err error) {
	ctx := context.Background()
	IAMService, err = iam.NewService(ctx)
	return
}
