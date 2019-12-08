package util

import (
	"context"
	"encoding/base64"
	"log"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iam/v1"
)

var iamService *iam.Service

func init() {
	var err error
	iamService, err = iam.NewService(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

type SignedURLType string

const (
	SignedURLTypeGET SignedURLType = "GET"
	SignedURLTypePUT SignedURLType = "PUT"
)

func GetSignedURL(ctx context.Context, method SignedURLType, bucketName, fileName, contentType string) (string, error) {
	return storage.SignedURL(bucketName, fileName, &storage.SignedURLOptions{
		GoogleAccessID: ServiceAccountName,
		Method:         string(method),
		Expires:        time.Now().Add(15 * time.Minute),
		ContentType:    contentTyp,
		SignBytes: func(b []byte) ([]byte, error) {
			resp, err := iamService.Projects.ServiceAccounts.SignBlob(
				ServiceAccountID,
				&iam.SignBlobRequest{BytesToSign: base64.StdEncoding.EncodeToString(b)},
			).Context(ctx).Do()
			if err != nil {
				return nil, err
			}
			return base64.StdEncoding.DecodeString(resp.Signature)
		},
	})
}
