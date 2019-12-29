package util

import (
	"context"
	"encoding/base64"
	"gae-test-project/connection"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iam/v1"
)

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
		ContentType:    contentType,
		SignBytes: func(b []byte) ([]byte, error) {
			resp, err := connection.IAMService.Projects.ServiceAccounts.SignBlob(
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
