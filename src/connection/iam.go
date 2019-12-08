package connection

import (
	"context"
	"log"

	"google.golang.org/api/iam/v1"
)

var IAMService *iam.Service

func init() {
	iamService, err := iam.NewService(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	IAMService = iamService
}
